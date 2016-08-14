package main

import (
	//	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "grpc_bs_benchmark/proto"

	"fmt"
	"io"
	"net"
	"runtime"
	"sync/atomic"
	"time"
)

type echoServiceServer struct{}

var num_connections int32
var num_packets int32
var max_connections int32

func (s *echoServiceServer) Echo(stream pb.EchoService_EchoServer) error {
	atomic.AddInt32(&num_connections, 1)
	defer func() {
		if err := recover(); err != nil {
			grpclog.Printf("failed: %v", err)
		}
		atomic.AddInt32(&num_connections, -1)
	}()

	if max_connections < num_connections {
		max_connections = num_connections
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}
		//grpclog.Printf("recv msg=%v, stream=%v", in.Msg, &stream)
		atomic.AddInt32(&num_packets, 1)

		var eh = new(pb.EchoHello)
		eh.Echomsg = fmt.Sprintf("you send msg: %s", in.Msg)
		if err := stream.Send(eh); err != nil {
			grpclog.Printf("send error")
			return err
		}
	}
}

func perf() {
	start := time.Now()
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("run (%v)seconds, max connections(%v), current_connections(%v), recv packets(%d)\n", time.Now().Sub(start).Seconds(), max_connections, num_connections, num_packets)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lis, err := net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%d", 6666))
	if err != nil {
		grpclog.Fatalf("failed to listen 6666: %v", err)
		return
	}
	var opts []grpc.ServerOption
	//	opts = append(opts, grpc.MaxConcurrentStreams(2000000))
	gServer := grpc.NewServer(opts...)
	s := new(echoServiceServer)
	pb.RegisterEchoServiceServer(gServer, s)
	go perf()
	gServer.Serve(lis)
}
