package main

import (
	"database"
	"log"
	"logger"
	"net/http"
	"router"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/teams", router.Route)
	logger.LogOut("Escuchando en localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	database.Dispose()
}
