package day23

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"aoc"
	"aoc/intcode"
)

type Address int

type Packet struct {
	Addr Address
	X, Y int
}

type Peer interface {
	Send(Packet)
	WhoAmI() Address
}

type Router struct {
	ingress         chan Packet
	peers           map[Address]Peer
	lastRXTimestamp int64
}

func NewRouter() *Router {
	c := &Router{
		ingress: make(chan Packet, 1024),
		peers:   make(map[Address]Peer, 64),
	}
	go c.Recv()
	return c
}

func (c *Router) WhoAmI() Address { return 1 << 31 }
func (c *Router) Send(p Packet)   { c.ingress <- p }
func (c *Router) Register(p Peer) { c.peers[p.WhoAmI()] = p }

func (c *Router) Recv() {
	for p := range c.ingress {
		atomic.StoreInt64(&c.lastRXTimestamp, time.Now().UnixNano())
		peer, ok := c.peers[p.Addr]
		if !ok {
			panic(fmt.Sprintf("no such peer: %d", p.Addr))
		}
		peer.Send(p)
	}
}

func (c *Router) LastRXTimestamp() (v int64) {
	return atomic.LoadInt64(&c.lastRXTimestamp)
}

type NAT struct {
	C chan int // send tests answers to caller

	addr   Address
	router *Router

	// the following are shared by the Send callers' and Poll goroutines.
	mu                 sync.Mutex
	lastPacketReceived Packet
	lastPacketSent     Packet
}

func NewNAT(addr Address, router *Router) *NAT {
	c := &NAT{
		C:      make(chan int, 1),
		addr:   addr,
		router: router,
	}
	go c.Poll()
	return c
}

func (c *NAT) Send(p Packet) {
	aoc.Debugf("[NAT] RECV: %v\n", p)
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lastPacketReceived == (Packet{}) {
		c.C <- p.Y // found the answer to part 1!
	}
	c.lastPacketReceived = p
}

func (c *NAT) WhoAmI() Address { return c.addr }

func (c *NAT) Poll() {
	idleTimeout := time.Millisecond * 100
	for {
		// the best signal I can see of an idle network is simply some delay since
		// any packet traversed the router
		time.Sleep(idleTimeout)
		now := time.Now().UnixNano()
		if now-c.router.LastRXTimestamp() < int64(idleTimeout) {
			continue
		}
		func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			p := c.lastPacketReceived
			p.Addr = 0 // rewrite to addr:0
			if p == c.lastPacketSent {
				c.C <- p.Y // found the answer to part 2!
			}
			aoc.Debugf("[NAT] SEND: %v\n", p)
			c.lastPacketSent = p
			c.router.Send(p)
		}()
	}
}

type NIC struct {
	addr Address

	// ingress is shared by the VM.Run and Recv goroutines
	ingress    chan Packet
	ingressMu  sync.Mutex
	ingressBuf []int

	// egress is owned by the VM.Run goroutine
	egress    Peer
	egressBuf []int
}

func NewNIC(addr Address, peer Peer) *NIC {
	nic := &NIC{
		addr:       addr,
		ingress:    make(chan Packet, 64),
		ingressBuf: make([]int, 1, 64), // 1 for init address
		egress:     peer,
		egressBuf:  make([]int, 0, 64),
	}
	nic.ingressBuf[0] = int(addr) // init address
	go nic.Recv()
	return nic
}

func (c *NIC) WhoAmI() Address { return c.addr }
func (c *NIC) Send(p Packet)   { c.ingress <- p }

// Recv dequeues packets, decodes and pushes them to the VM's input buffer.
func (c *NIC) Recv() {
	for p := range c.ingress {
		aoc.Debugf("[%d] RECV: %v\n", c.addr, p)
		func() {
			c.ingressMu.Lock()
			defer c.ingressMu.Unlock()
			c.ingressBuf = append(c.ingressBuf, p.X, p.Y)
		}()
	}
}

// ReadInt is called by the VM to receive packet data.
func (c *NIC) ReadInt() (v int, err error) {
	c.ingressMu.Lock()
	defer c.ingressMu.Unlock()
	if len(c.ingressBuf) == 0 {
		time.Sleep(time.Millisecond * 10) // idle-wait backoff
		return -1, nil                    // no packets in buffer
	}
	v = c.ingressBuf[0]
	c.ingressBuf = c.ingressBuf[1:]
	return
}

// WriteInt is called by the VM to send packets.
func (c *NIC) WriteInt(v int) (err error) {
	c.egressBuf = append(c.egressBuf, v)
	if len(c.egressBuf) < 3 {
		return
	}
	if len(c.egressBuf) > 3 {
		panic(len(c.egressBuf))
	}
	p := Packet{
		Addr: Address(c.egressBuf[0]),
		X:    c.egressBuf[1],
		Y:    c.egressBuf[2],
	}
	c.egressBuf = c.egressBuf[:0]
	aoc.Debugf("[%d] SEND: %v\n", c.addr, p)
	c.egress.Send(p)
	return
}

type Computer struct {
	vm intcode.VM
}

func NewComputer(nic *NIC) *Computer {
	data, err := intcode.OpenData(aoc.Fixture("day23"))
	if err != nil {
		panic(err)
	}
	return &Computer{
		vm: intcode.NewWithIO(data, nic, nic),
	}
}

func (c *Computer) Run() (err error) {
	_, err = c.vm.Run()
	return
}
