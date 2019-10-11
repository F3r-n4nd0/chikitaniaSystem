package eventstore

import (
	"eventStore/protobuf"
	"google.golang.org/grpc"
	"net"
)

type gRpcServer struct {
	network string
	address   string
}

func NewGRpcServer(address string) *gRpcServer  {
	return &gRpcServer{
		network: "tcp",
		address:   address,
	}
}

func (g gRpcServer) Listen(srv protobuf.EventStoreServer) error {
	lis, err := net.Listen(g.network, g.address)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	protobuf.RegisterEventStoreServer(s, srv)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}