package main

import (
	"os"
	"os/signal"
	"syscall"

	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "grpc_bs_benchmark/proto"
	"net"
	"sync"
	//"sync/atomic"
	"time"
)

var (
	max_clients = 30000
	max_packtes = 10000
)

func mydial(addr string, t time.Duration) (net.Conn, error) {
	var ief, _ = net.InterfaceByName("lo")
	var addrs, _ = ief.Addrs()
	//fmt.Println("net addr=", addrs[0])
	//fmt.Println("remote addr=", addr)
	var netaddr = &net.TCPAddr{
		IP: addrs[0].(*net.IPNet).IP,
	}
	dialer := net.Dialer{LocalAddr: netaddr}
	return dialer.Dial("tcp4", addr)
}

func client_stream(wg *sync.WaitGroup, chann chan int) {
	defer wg.Done()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithDialer(mydial))
	conn, err := grpc.Dial("localhost:6666", opts...)
	if err != nil {
		grpclog.Fatalf("failed to dial server")
		return
	}
	defer conn.Close()
	client := pb.NewEchoServiceClient(conn)

	stream, err := client.Echo(context.Background())
	if err != nil {
		grpclog.Fatalf("get stream error")
	}
	for i := 0; i < max_packtes; i++ {
		// send
		h := new(pb.Hello)
		h.Msg = "hello grpc"
		err = stream.Send(h)
		if err != nil {
			grpclog.Fatalf("send erroor")
			return
		}
		// recv
		_, err := stream.Recv()
		if err != nil {
			grpclog.Fatalf("recv error")
		}

		// break or sleep
		close := false
		select {
		case <-chann:
			close = true
		default:
			time.Sleep(5 * time.Second)
		}
		if close {
			stream.CloseSend()
			break
		}
	}
}

func StartSignal(channs []chan int) {
	var (
		c chan os.Signal
		s os.Signal
	)
	c = make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,
		syscall.SIGINT, syscall.SIGSTOP)
	// Block until a signal is received.
	for {
		s = <-c
		fmt.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			for i := 0; i < max_clients; i++ {
				close(channs[i])
			}
			return
		case syscall.SIGHUP:
			// TODO reload
			//return
		default:
			fmt.Println("unknow signal:", s.String())
			return
		}
	}

}

func main() {
	var channs []chan int
	for i := 0; i < max_clients; i++ {
		chann := make(chan int)
		channs = append(channs, chann)
	}
	// catch term signal
	go StartSignal(channs)
	var wgs []*sync.WaitGroup
	for i := 0; i < max_clients; i++ {
		var wg = new(sync.WaitGroup)
		wg.Add(1)
		go client_stream(wg, channs[i])
		wgs = append(wgs, wg)
	}

	// wait for go routine return
	for _, wg := range wgs {
		wg.Wait()
	}
}
