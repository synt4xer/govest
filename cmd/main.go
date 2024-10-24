package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/synt4xer/govest/config"
	"github.com/synt4xer/govest/internal/database"
	"github.com/synt4xer/govest/internal/gql"
)

func main() {
	ctx := context.Background()

	cfg, err := config.ProvideConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.ProvideDB(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}}))

	r := chi.NewRouter()
	if cfg.Server.GoEnv == "development" {
		r.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost%s/ for GraphQL playground", cfg.Server.Port)
	}
	r.Handle("/query", srv)

	port := cfg.Server.Port
	log.Fatal(http.ListenAndServe(port, r))
}
