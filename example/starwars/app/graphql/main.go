package main

import (
	"log"
	"net/http"

	"github.com/jho3r/grpc-graphql-gateway/example/starwars/spec/starwars"
	"github.com/jho3r/grpc-graphql-gateway/runtime"
)

func main() {
	mux := runtime.NewServeMux(runtime.Cors())

	if err := starwars.RegisterStartwarsServiceGraphqlHandler(mux, nil); err != nil {
		log.Fatalln(err)
	}
	http.Handle("/graphql", mux)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}
