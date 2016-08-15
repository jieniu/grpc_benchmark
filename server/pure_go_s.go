package main

import (
	"fmt"
	"net"
	"runtime"
	"sync/atomic"
	"time"
)

var num_conns int32

func outputinfo() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("num_conns=%d\n", num_conns)

	}
}

func handle(c net.Conn) error {
	defer c.Close()
	atomic.AddInt32(&num_conns, 1)
	defer atomic.AddInt32(&num_conns, -1)
	var buf = make([]byte, 10)
	var n int
	var err error
	for {
		// read
		n, err = c.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			return err
		}
		// write
		n, err = c.Write(buf[:n])
		if err != nil {
			fmt.Println("write error:", err)
			return err
		}

	}
	return nil
}

func acception(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Printf("accept failed: %v\n", err)
			continue
		}

		go handle(c)
	}

}

func main() {
	num_conns = 0
	runtime.GOMAXPROCS(runtime.NumCPU())
	go outputinfo()
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Printf("listen failed: %v\n", err)
		return
	}
	l2, err := net.Listen("tcp", ":6667")
	if err != nil {
		fmt.Printf("listen failed: %v\n", err)
		return
	}
	defer l.Close()
	defer l2.Close()
	go acception(l)
	go acception(l2)
	for {
		time.Sleep(2 * time.Second)
	}
}
