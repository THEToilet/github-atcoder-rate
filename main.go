package main

import (
	"github-program-rate/pkg/gateway"
	"log"
)

func main() {
	server := gateway.NewServer(logger)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
