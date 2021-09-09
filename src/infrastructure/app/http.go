package app

import (
	"log"
	"net/http"
	"time"

	"github.com/Dall06/go-cleanarch-template/config"
	s "github.com/Dall06/go-cleanarch-template/src/infrastructure/server"
	"github.com/gorilla/mux"
)


func StartHTTPServer(router *mux.Router) {
	server := &http.Server{
		Handler: router,
		Addr:    config.Port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server has started at port", config.Port)
	s.NewGracefullyShutDown(router, config.Port).RunGracefully()
	log.Fatal(server.ListenAndServe())
}
