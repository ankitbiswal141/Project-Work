package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct{
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL string `envconfig:"ORDER_SERVICE_URL"`
}

func main(){
	var cfg AppConfig
	err := envconfig.Process("",&cfg)

	if err != nil {
		log.Fatal(err)
	}

	server,err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL)

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql",handler.GraphQL((server.ToExecutableSchema())))
	http.Handle("/playground", playground.Handler("ankit","/graphql"))
}