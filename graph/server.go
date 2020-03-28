package graph

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aq1018/gogqlpoc/db"
	"github.com/aq1018/gogqlpoc/graph/generated"
	"github.com/aq1018/gogqlpoc/operation"
)

type Server struct {
	op  *operation.Operation
	srv http.Handler
}

func NewServer() *Server {
	op := operation.NewOperation()
	middleware := operation.Middleware(op)

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}})
	srv := handler.NewDefaultServer(schema)

	return &Server{srv: middleware(srv), op: op}
}

func (server *Server) Run(port string) {
	db.AutoMigrate(server.op.DB)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server.srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
