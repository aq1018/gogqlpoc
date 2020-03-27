package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PeerStreet/aqgqlpoc/graph"
	"github.com/PeerStreet/aqgqlpoc/graph/generated"
	"github.com/PeerStreet/aqgqlpoc/graph/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const defaultPort = "8080"

func main() {
	db, err := gorm.Open("postgres", os.Getenv("DBURL"))
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	db.AutoMigrate(&model.Loan{}, &model.Property{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
