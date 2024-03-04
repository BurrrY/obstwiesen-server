package main

import (
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BurrrY/obstwiesen-server/graph"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000", "http://192.168.178.201:3000"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "localhost"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	fs := http.FileServer(http.Dir(os.Getenv("FILE_CONNSTR")))
	router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	router.Handle("/", playground.Handler("Obstwiese", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
func main2() {
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
