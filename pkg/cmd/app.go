package cmd

import (
	"log"

	"github.com/Dall06/go-cleanarch-template/pkg/database"
	"github.com/Dall06/go-cleanarch-template/pkg/cmd/rest"
	"github.com/gorilla/mux"
)

func RunApp() {
	router := mux.NewRouter()

	mysqlConn := database.NewMySQLConn()
	conn, err := mysqlConn.OpenConnection()
	if err != nil {
		log.Fatalf("ERROR %v", err)
	}

	httpServer := rest.NewHTTPServer(router, conn)
	httpServer.StartHTTPServer()
}
