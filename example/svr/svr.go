package main

import (
	"context"
	"flag"
	"github.com/sereiner/library/balancer"
	logger "github.com/sereiner/library/log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	pb "github.com/sereiner/parrot/example/helloworld"
)

var (
	log  = logger.GetSession("grpc", logger.CreateSession())
	serv = flag.String("service", "hello_service", "service name")
	host = flag.String("host", "localhost", "listening host")
	port = flag.String("port", "50001", "listening port")
	reg  = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", net.JoinHostPort(*host, *port))
	if err != nil {
		panic(err)
	}

	err = balancer.Register(*reg, *serv, *host, *port, time.Second*10, 15)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Infof("receive signal '%v'", s)
		balancer.UnRegister()
		os.Exit(1)
	}()

	log.Infof("starting hello service at %s", *port)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Infof(" Receive is %s", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name + " from " + net.JoinHostPort(*host, *port)}, nil
}
