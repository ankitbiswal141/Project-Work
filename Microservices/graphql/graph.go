package main

import (
	"github.com/99designs/gqlgen/graphql"
	"golang.org/x/text/message/catalog"
)

type Server struct{
  accountClient *account.Client
  catalogClient *catalog.Client
  orderClient   *order.Client
}

// It takes three queries for the database client
func NewGraphQLServer(accountUrl, catalogUrl,orderUrl string) (*Server,error){
  accountClient, err := account.NewClient(accountUrl)
  
  if err != nil {
    return nil,err
  }

  catalogClient, err := catalog.NewClient(catalogUrl)

  if err != nil{
    accountClient.Close()
    return nil,err
  }

  orderClient, err := order.NewClient(orderUrl)

  if err != nil {
    accountClient.Close()
    catalogClient.Close()
    return nil, err
  }

  return &Server{
    accountClient,
    catalogClient,
    orderClient,
  },nil
}

func (s *Server) Mutation() MutationResolver {
  return &mutationResolver{
    server: s,
  }
}

func (s *Server) Query() QueryResolver{
  return &queryResolver{
    server: s,
  }
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema{
  return NewExecutableSchema(Config{
    Resolvers: s, 
  })
}