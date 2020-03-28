package main

import (
	"os"

	"github.com/aq1018/gogqlpoc/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	graph.NewServer().Run(port)
}
