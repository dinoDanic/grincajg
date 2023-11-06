package main

import (
	"grincajg/database"
	"grincajg/database/seeds"
	"grincajg/env"
	"grincajg/graph"
	"grincajg/middleware"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

const defaultPort = ":8090"

func main() {
	env.LoadEnv()
	database.Connect()
	seeds.SeedAll()

	router := chi.NewRouter()

	router.Use(middleware.Middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("Grincajg", "/api"))
	router.Handle("/api", srv)

	err := http.ListenAndServe(defaultPort, router)

	if err != nil {
		panic(err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(defaultPort, nil))

}
