package intcode

type SyncIO interface {
	IntReadWriter

	Run() (v int, err error)
}

type syncIO struct {
	vm     VM
	stdin  IntBuffer
	stdout IntBuffer
}

// NewSyncIO returns a virtual machine that performs blocking IO, running the
// VM program until the desired IO is available.
func NewSyncIO(data Data) SyncIO {
	stdin := NewIntBuffer(nil)
	stdout := NewIntBuffer(nil)
	vm := NewWithIO(data, stdin, stdout)
	return &syncIO{
		stdin:  stdin,
		stdout: stdout,
		vm:     vm,
	}
}

// Read performs a blocking read of a single value from the VM's stdout device.
func (c *syncIO) ReadInt() (v int, err error) {
	for c.stdout.Len() == 0 {
		err = c.vm.Step()
		if err != nil {
			return
		}
	}
	return c.stdout.ReadInt()
}

// Write performs a blocking write of a single value to the VM's stdin device.
func (c *syncIO) WriteInt(v int) (err error) {
	err = c.stdin.WriteInt(v)
	if err != nil {
		return
	}
	for c.stdin.Len() != 0 {
		err = c.vm.Step()
		if err != nil {
			return
		}
	}
	return
}

// Run the VM until it returns an error.
func (c *syncIO) Run() (v int, err error) {
	return c.vm.Run()
}
