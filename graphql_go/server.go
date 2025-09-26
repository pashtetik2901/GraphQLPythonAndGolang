package main

import (
	"log"
	"net/http"
	"os"

	"graphql_go/database"
	"graphql_go/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Создаем резолвер с передачей DB
	resolver := &graph.Resolver{
		DB: database.DB,
	}

	// Создаем GraphQL handler с правильными настройками
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))


	http.Handle("/query", srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


