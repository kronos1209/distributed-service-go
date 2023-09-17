package main

import (
	"log"

	"github.com/kronos1209/proglog/internal/server"
)

func main() {
	src := server.NewHTTPServer(":8080")
	log.Fatal(src.ListenAndServe())
}
