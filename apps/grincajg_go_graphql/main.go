package main

import (
	"grincajg/database"
	"grincajg/env"
	"grincajg/graph"
	auth "grincajg/middleware"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	env.LoadEnv()
	database.Connect()

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("Grincajg", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":8090", router)
	if err != nil {
		panic(err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8090", nil))

}
