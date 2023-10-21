package main

import (
	"grincajg/database"
	"grincajg/env"
	"grincajg/graph"
	"grincajg/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	config := loadEnv()
	loadDatabase(config)
	serveApplication()
}

func loadEnv() env.Config {
	config, err := env.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	return config
}

func loadDatabase(config env.Config) {
	database.Connect(config)
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.Organization{})
}

func serveApplication() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
