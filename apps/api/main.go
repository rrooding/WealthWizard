package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"

	"wealth-wizard/api/database"
	"wealth-wizard/api/graph"
	"wealth-wizard/api/resolvers"
)

const defaultPort = "8080"

func databaseDSN() (string, error) {
	pgHost := os.Getenv("PGHOST")
	pgUser := os.Getenv("PGUSER")
	pgPassword := os.Getenv("PGPASSWORD")
	pgDatabase := os.Getenv("PGDATABASE")

	config := &database.Config{
		Host:     pgHost,
		User:     pgUser,
		Password: pgPassword,
		Database: pgDatabase,
	}

	return config.GetDSN()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dsn, err := databaseDSN()
	if err != nil {
		log.Fatal(err)
	}

	db := database.InitDB(dsn)
	defer func() { database.CloseDB(db) }()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
