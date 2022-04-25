package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:18888")
	if err != nil {
		log.Fatal(err)
	}
	err = run(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func run(conn net.Conn) error {
	// errstring-start
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)

		// ソケットがクローズされていることを、エラーの文字列から判断する
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			return err
		}

		handleRead(buf)
	}
	// errstring-end
	return nil
}

func handleRead(buf []byte) {}
