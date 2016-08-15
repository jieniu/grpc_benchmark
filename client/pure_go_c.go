package main

import (
	"fmt"
	"net"
	"time"
)

var max = 100000

func dial_server(port int) (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("10.10.191.7:%d", port))
	if err != nil {
		fmt.Printf("connect failed: %v\n", err)
		return nil, err
	}

	return conn, err
}

func handle(conn net.Conn) {
	defer conn.Close()
	var data []byte
	data = []byte("hello")
	var err error
	var buf = make([]byte, 10)
	for {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("write failed:", err)
			return
		}
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn read error:", err)
			return
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	var conns []net.Conn
	var ports []int
	ports = append(ports, 6666)
	ports = append(ports, 6667)
	for i := 0; i < max; i++ {
		port := ports[i%2]
		conn, err := dial_server(port)
		if err != nil {
			continue

		}
		conns = append(conns, conn)

	}
	for _, conn := range conns {
		go handle(conn)

	}
	for {
		time.Sleep(5000 * time.Second)
	}

}
