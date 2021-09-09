package app

import (
	"log"

	"github.com/Dall06/go-cleanarch-template/src/database"
	"github.com/gorilla/mux"
)

func RunApp() {
	router := mux.NewRouter()
	mysqlConn := database.NewMySQLConn()

	conn, err := mysqlConn.OpenConnection()
	if err != nil {
		log.Fatalf("ERROR %v", err)
	}

	StartRouter(router, conn)
	StartHTTPServer(router)
}
