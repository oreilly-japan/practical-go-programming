package main

import (
	"net"

	"github.com/Songmu/retry"
)

func main() {
	// error-retry-start
	var b []byte
	err := retry.Retry(2, 0, func() error {
		_, ierr := tcpClient.Read(b)
		return ierr
	})
	if err != nil {
		// error handling
	}
	// error-retry-end
}

var tcpClient *net.TCPConn
