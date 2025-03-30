package account

import (
	"context"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
)

type grpcServer struct{
	service Service
}

func ListenServer(s Service,port int) error{
	listen,err := net.Listen("tcp",fmt.Sprintf(":%d",port))
	if err!= nil{
		return err
	}

	server:= grpc.newServer()
	reflection.Register(server)

	return server.Serve(listen)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.) (*pb.,error) {
	a,err := s.service.PostAccount((ctx,r.Name)
	if err!=nil {
		return nil,err
	}

	return &pb.{}
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.) (*pb.,error) {
	a,err := s.service.GetAccount(ctx,r.ID)
	if err!= nil{
		return nil,err
	}

	return &pb.{}
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.) (*pb.,error) {
	a,err := s.service.GetAccounts(ctx,r.ID)
	if err!= nil {
		return nil,err
	}
	return &pb.{}
}