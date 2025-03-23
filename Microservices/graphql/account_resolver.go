package main

import "context"

type accountResolver struct{
	server *Server
}

func (r *accountResolver) Orders(ctx context.Context,acc *Account) ([] *Order,error){
	
}