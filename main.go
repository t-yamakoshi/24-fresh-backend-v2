package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-sql-driver/mysql"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/gqlgen"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/gqlgen/resolver"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/repository"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/usecase"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dsn := "root:@tcp(localhost:23306)/"
	cfg, err := NewConfig(dsn)
	if err != nil {
		log.Fatalf("failed creating mysql config: %v", err)
	}
	cfg.ParseTime = true
	cfg.DBName = "24-fresh"
	client, err := NewClient(cfg)

	userRepository := repository.NewUserRepository(
		client,
	)
	userUsecase := usecase.NewUserUsecase(
		userRepository,
	)

	resolver := resolver.Resolver{
		UserUsecase: userUsecase,
	}

	srv := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func NewConfig(dsn string) (*mysql.Config, error){
	if dsn == "" {
		return nil, fmt.Errorf("dsn is empty")
	}
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewClient(cfg *mysql.Config) (*entgen.Client, error){
	client, err := entgen.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return client, nil
}
