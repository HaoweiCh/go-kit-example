package handlers

import (
	"context"
	"strings"

	pb "go-kit-example/services/echo"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.EchoServer {
	return echoService{}
}

type echoService struct{}

// Echo implements Service.
func (s echoService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp = pb.EchoResponse{
		Out: in.In,
	}
	return &resp, nil
}

// Louder implements Service.
func (s echoService) Louder(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp = pb.EchoResponse{
		Out: strings.Repeat(in.In, int(in.Loudness)),
	}
	return &resp, nil
}

// LouderGet implements Service.
func (s echoService) LouderGet(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp = pb.EchoResponse{
		Out: strings.Repeat(in.In, int(in.Loudness)),
	}
	return &resp, nil
}
