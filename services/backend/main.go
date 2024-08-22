package main

import (
	"fmt"
	"log"
	"net/http"
	//"crypto/tls"
	"database/sql"
	
	_ "github.com/lib/pq"
	
	//TranspLayerCustom "students/tls"
	"students/db"
	"students/env"
	"students/api/router"
	
)

func main() {

	envm := env.SetupEnv("env/.env")
	fmt.Printf("Environment setup complete: %+v\n", envm)

	dbase, err := sql.Open(envm.DbDriver, envm.DbSource)
	if err != nil {
		log.Printf("failed to open the database connection: %v\n", err)
	}

	err = dbase.Ping()
	if err != nil {
		dbase.Close()
		log.Printf("failed to ping the database: %v\n", err)
	}
	
	slqcDatabase := db.New(dbase)
	fmt.Printf("Database Connected\n")
/* 
	config := TranspLayerCustom.NewConfig()
	fmt.Println("TLS configuration set up") */

	router := router.NewRouter(slqcDatabase)
	fmt.Println("Router setup complete")

	/* fmt.Println("Creating and configuring the HTTP server")
	s := &http.Server{
		Addr:         ":8443",
		Handler:      router,
		TLSConfig:    config,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	} */

	fmt.Printf("Listening on port: 8443\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
