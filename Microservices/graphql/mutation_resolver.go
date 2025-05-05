package main

import "context"

type mutationResolver struct{
	server *Server
}

func (r *mutationResolver) createAccout(ctx context.Context, in AccountInput) (*Account,error){

}

func (r *mutationResolver) createProduct(ctx context.Context, in ProductInput) (*Account,error){
	
}

func (r *mutationResolver) createOrder(ctx context.Context, in OrderInput) (*Account,error){
	
}