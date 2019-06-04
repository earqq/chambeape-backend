package main

import (
	"chambeape/db"
	"chambeape/graphql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "8085"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db.ConnectDB()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()

	r.Handle("/", handler.Playground("User", "/graphql"))
	r.Handle("/graphql", c.Handler(handler.GraphQL(graphql.NewExecutableSchema(graphql.New()),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}))),
	)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8085", nil))
}
