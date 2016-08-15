package main

import (
	"os"
	"os/signal"
	"syscall"

	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "grpc_bs_benchmark/proto"
	"math/rand"
	"net"
	"sync"
	"time"
)

var (
	max_clients = 60000
	max_packtes = 10000
	Index       int
)

func init() {
	flag.IntVar(&Index, "interface", 0, "assign net interface index")
}

func mydial(addr string, t time.Duration) (net.Conn, error) {
	var ief, _ = net.InterfaceByName("eth0")
	var addrs, _ = ief.Addrs()
	var netaddr = &net.TCPAddr{
		IP: addrs[Index].(*net.IPNet).IP,
	}
	dialer := net.Dialer{LocalAddr: netaddr}
	return dialer.Dial("tcp4", addr)
}

func client_stream(wg *sync.WaitGroup, chann chan int) {
	defer wg.Done()
	var opts []grpc.DialOption
	var h = new(pb.Hello)
	var err error
	var close bool
	var conn *grpc.ClientConn
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithDialer(mydial))
	conn, err = grpc.Dial("10.10.191.7:6666", opts...)
	if err != nil {
		grpclog.Printf("failed to dial server, %v\n", err)
		return
	}
	defer conn.Close()
	client := pb.NewEchoServiceClient(conn)

	stream, err := client.Echo(context.Background())
	if err != nil {
		grpclog.Printf("get stream error %s\n", err)
		return
	}

	var sec time.Duration
	sec = (time.Duration)(rand.Intn(10) + 1)
	time.Sleep(sec * time.Second)

	for i := 0; i < max_packtes; i++ {
		// send

		h.Msg = "hello grpc"
		err = stream.Send(h)
		if err != nil {
			grpclog.Printf("send erroor, %v\n", err)
			return
		}
		// recv
		_, err = stream.Recv()
		if err != nil {
			grpclog.Printf("recv error, %v\n", err)
			return
		}

		// break or sleep
		close = false
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
	flag.Parse()

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
