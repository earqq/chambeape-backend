package main

import (
	"log"
	"net/http"
	"os"
	"tuchamba/db"
	"tuchamba/graphql"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "8083"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db.ConnectDB()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:" + port},
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

	log.Fatal(http.ListenAndServe(":8083", nil))
}
