package day23

import (
	"aoc"
	"testing"
)

func TestDay23(t *testing.T) {
	router := NewRouter()

	// NAT device will wake the idle network and watch for duplicate packets to
	// address [0]
	nat := NewNAT(255, router)
	router.Register(nat)

	// build 50x NICS to send and receive packets through the router
	nics := make([]*NIC, 50)
	for i := 0; i < len(nics); i++ {
		nics[i] = NewNIC(Address(i), router)
		router.Register(nics[i])
	}

	// attach to VMs and start them up
	for i := 0; i < len(nics); i++ {
		c := NewComputer(nics[i])
		go c.Run()
	}

	// get Packet.Y of first packet sent to NAT
	t.Run("Part1", func(t *testing.T) {
		v := <-nat.C
		aoc.AssertInt(t, 22650, v, "bad Packet.Y value")
	})

	// get Packet.Y of first duplicate packet sent by the NAT
	t.Run("Part2", func(t *testing.T) {
		v := <-nat.C
		aoc.AssertInt(t, 17298, v, "bad Packet.Y value")
	})
}
