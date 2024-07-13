package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"students/api/router"
	"students/db"
	"students/env"
	tls_ "students/tls"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Setting up environment variables from ./env/.env")
	envm := env.SetupEnv("./env/.env")
	fmt.Printf("Environment setup complete: %+v\n", envm)

	fmt.Printf("Opening the database\n")
	dbase, err := sql.Open(envm.DbDriver, envm.DbSource)
	if err != nil {
		log.Printf("failed to open the database connection: %v\n", err)
	}

	fmt.Printf("Pinging Database\n")
	err = dbase.Ping()
	if err != nil {
		dbase.Close()
		log.Printf("failed to ping the database: %v\n", err)
	}
	fmt.Printf("Database Connected\n")

	fmt.Println("Initializing database connection")
	slqcDatabase := db.New(dbase)
	fmt.Println("Database connection initialized")

	fmt.Println("Setting up TLS configuration")
	config := tls_.NewConfig()
	fmt.Println("TLS configuration set up")

	fmt.Println("Setting up router")
	router := router.NewRouter(slqcDatabase)
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
