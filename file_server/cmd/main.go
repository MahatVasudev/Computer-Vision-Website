package main

import (
	"log"

	"github.com/MahatVasudev/liveStreamingProject/file_server/cmd/api"
)

func main() {

	addr := ":6000"

	server := api.NewApiServer(addr)

	if err := server.Run(); err != nil {
		log.Fatalf("Error Occurred: %v", err)
	}
}
