package account

import (
	"context"
	"fmt"
	"net"

	"github.com/ankitbiswal141/Project-Work/Microservices/account/pb/github.com/ankitbiswal141/Project-Work/Microservices/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func ListenServer(s Service, port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &grpcServer{s})
	reflection.Register(server)

	return server.Serve(listen)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	res, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}

	return &pb.PostAccountResponse{
		Account: &pb.Account{
			Id:   res.ID,
			Name: res.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	res, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   res.ID,
			Name: res.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	res, err := s.service.GetAccounts(ctx, r.)
	if err != nil {
		return nil, err
	}

	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(accounts,
			&pb.Account{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}

	return &pb.GetAccountsResponse{
		Accounts: accounts,
	}, nil
}
