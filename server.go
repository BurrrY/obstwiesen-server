package main

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BurrrY/obstwiesen-server/graph"
	"github.com/BurrrY/obstwiesen-server/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

const defaultPort = "8080"

func setup() {
	log.Info("Init Main")
	viper.AutomaticEnv()
	viper.SetDefault(config.DB_PROVIDER, "hehe")
	viper.SetDefault(config.DB_CONNSTR, "")
	viper.SetDefault(config.DB_NAME, "meadow")

	viper.SetDefault(config.FILE_CONNSTR, "./files")
	viper.SetDefault(config.FILE_PROVIDER, "disk")

	viper.SetDefault(config.PUBLIC_URL, "localhost:8080")

	viper.SetDefault(config.GQL_PORT, "8080")
	viper.SetDefault(config.GQL_PATH, "/graphql")

}

func main() {
	setup()

	b, _ := json.MarshalIndent(viper.AllSettings(), "", "  ")
	fmt.Print(string(b), "\n")

	r := graph.Resolver{}
	r.Setup()

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

	fs := http.FileServer(http.Dir(viper.GetString(config.FILE_CONNSTR)))
	router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	router.Handle("/", playground.Handler("Obstwiese", viper.GetString(config.GQL_PATH)))
	router.Handle(viper.GetString(config.GQL_PATH), srv)

	err := http.ListenAndServe(":"+viper.GetString(config.GQL_PORT), router)
	if err != nil {
		panic(err)
	}

}
