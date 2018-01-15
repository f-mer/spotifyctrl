package portscanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func IsOpen(port int) bool {
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	conn, err := net.DialTimeout("tcp", addr, time.Duration(500)*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

func FirstOpened(start, end int) (int, bool) {
	ch := make(chan int)
	defer close(ch)
	done := make(chan struct{})

	var wg sync.WaitGroup
	for port := start; port <= end; port++ {
		wg.Add(1)
		go func(port int) {
			if IsOpen(port) {
				ch <- port
			}
			wg.Done()
		}(port)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case port := <-ch:
		return port, true
	case <-done:
		return 0, false
	}
}
