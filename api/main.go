package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello there.")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found")
	}

	fmt.Println("current port:", portString)

	// define router
	router := chi.NewRouter()

	// cors config for router, let user access from browser
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           600,
	}))

	v1Router := chi.NewRouter()
	// hook up handlerReadiness function to "/ready" path via GET request
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerError)

	// full path: /v1/ready
	router.Mount("/v1", v1Router)

	// define server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
