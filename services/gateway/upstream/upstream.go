package upstream

import "google.golang.org/grpc"

type GRPCUpstream struct {
	address string
	conn    *grpc.ClientConn
	opts    []grpc.DialOption
}

func NewGRPCUpstream(address string, opts ...grpc.DialOption) (*GRPCUpstream, error) {
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	return &GRPCUpstream{
		address: address,
		conn:    conn,
		opts:    opts,
	}, nil
}

type UpstreamManager struct {
	upstreams map[string]*GRPCUpstream
}

func NewUpstreamManager() *UpstreamManager {
	return &UpstreamManager{
		upstreams: make(map[string]*GRPCUpstream),
	}
}
