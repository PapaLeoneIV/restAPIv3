package main

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"students/db"
	"students/env"
	"students/api/router"
	tls_ "students/tls"
)

func main() {
	fmt.Println("Setting up environment variables from env/.env")
	envManager := env.SetupEnv("env/.env")
	fmt.Printf("Environment setup complete: %+v\n", envManager)

	fmt.Println("Initializing database connection")
	db := db.NewDB(envManager.DbDriver, envManager.DbSource)
	fmt.Println("Database connection initialized")

	fmt.Println("Setting up TLS configuration")
	config := tls_.NewConfig()
	fmt.Println("TLS configuration set up") 

	fmt.Println("Setting up router")
	router := router.NewRouter(db)
	fmt.Println("Router setup complete")

	fmt.Println("Creating and configuring the HTTP server")
	s := &http.Server{
		Addr:         ":8443",
		Handler:      router,
		 TLSConfig:    config,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0), 
	}

	fmt.Printf("Listening on port: 8443\n")
	if err := s.ListenAndServeTLS("tls/server.crt", "tls/server.key"); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}