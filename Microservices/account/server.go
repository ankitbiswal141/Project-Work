package account

import (
	"context"
	"net"
	"fmt"
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

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse,error) {
	res,err := s.service.PostAccount(ctx,r.Name)
	if err!=nil {
		return nil,err
	}

	return &pb.PostAccountResponse{
		Account: &pb.Account{
			Id: res.ID,
			Name: res.Name,
		},
	},nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse,error) {
	res,err := s.service.GetAccount(ctx,r.ID)
	if err!= nil{
		return nil,err
	}

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			ID: res.ID,
			Name: res.Name,
		},
	},nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse,error) {
	res,err := s.service.GetAccounts(ctx,r.ID)
	if err!= nil {
		return nil,err
	}

	accounts := []*pb.Account{}
	for _,p := range res {
		accounts = append(accounts, 
			&pb.Account{
				ID: p.ID,
				Name: p.Name,
			},
		)
	}

	return &pb.GetAccountsResponse{
		Accounts:accounts,
	},nil
}