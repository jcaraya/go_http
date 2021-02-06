package memory

import "time"

type bigInt struct {
	a [100000]int
}

// AllocateMemory TODO:
func AllocateMemory(timeStep time.Duration) {
	allocated := []bigInt{}
	for i := 0; i < 10; i++ {
		allocated = append(allocated, bigInt{})
		time.Sleep(timeStep)
	}

	for {
	}
}
