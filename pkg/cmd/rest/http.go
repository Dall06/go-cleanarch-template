package rest

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Dall06/go-cleanarch-template/config"
	"github.com/Dall06/go-cleanarch-template/pkg/infrastructure/services"
	"github.com/Dall06/go-cleanarch-template/pkg/infrastructure/middleware"
	"github.com/Dall06/go-cleanarch-template/pkg/infrastructure/routes"

	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user/delivery/handler"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user/repository/mysqldb"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user/usecase"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Router *mux.Router
	Conn *sql.DB
}

func NewHTTPServer(router *mux.Router, conn *sql.DB) *httpServer{
	return &httpServer{
		Router: router,
		Conn: conn,
	}
}

func (httpS *httpServer) setRouter() {
	userMysqlRepo := mysqldb.MySQLUserRepository(httpS.Conn)
	userInteractor := usecase.NewUserInteractor(userMysqlRepo)
	userHandler := handler.NewUserHandler(userInteractor)

	amw := middleware.NewAuthenticationMiddleWre()
	amw.Populate()

	cors := middleware.NewCORSMiddleware("GET, POST, PUT, DELETE", "localhost")
	
	routes.NewUserRoutes(httpS.Router, userHandler).SetUserRoutes()

	httpS.Router.Use(amw.Middleware)
	httpS.Router.Use(cors.EnableCORS)
}


func (httpS *httpServer) StartHTTPServer() {
	server := &http.Server{
		Handler: httpS.Router,
		Addr:    config.Port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	httpS.setRouter()

	log.Println("Server has started at port", config.Port)
	services.NewGracefullyShutDown(httpS.Router, config.Port).RunGracefully()
	log.Fatal(server.ListenAndServe())
}
