package main

import (
	"log"

	server "github.com/aidk/gocatfacts/cmd/internal/api/http"
	"github.com/aidk/gocatfacts/cmd/internal/api/service"
	"github.com/aidk/gocatfacts/cmd/internal/logging"
)

func main() {

	// we initialize our fact service
	svc := service.NewFactService("https://catfact.ninja/fact")

	// we wrap our fact service with a logging service
	svc = logging.NewLoggingService(svc)

	// we initialize our server
	api := server.NewServer(svc)

	// we start the server and return any error that occurs if the server fails to start
	log.Fatal(api.Start(":3000"))
}
