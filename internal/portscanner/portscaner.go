package portscanner

import (
	"fmt"
	"net"
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

func FirstOpened(start, end int) int {
	ch := make(chan int)
	for port := start; port <= end; port++ {
		go func(port int) {
			if IsOpen(port) {
				ch <- port
			}
		}(port)
	}
	return <-ch
}
