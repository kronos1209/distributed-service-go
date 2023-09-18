package main

import (
	"log"

	"github.com/kronos1209/proglog/internal/server"
)

func main() {
	src, _ := server.NewGRPCServer(nil)
	log.Fatal(src)
}
