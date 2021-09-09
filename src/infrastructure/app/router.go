package app

import (
	"database/sql"

	"github.com/Dall06/go-cleanarch-template/config"
	"github.com/Dall06/go-cleanarch-template/src/infrastructure/middleware"
	"github.com/Dall06/go-cleanarch-template/src/infrastructure/routes"
	"github.com/Dall06/go-cleanarch-template/src/infrastructure/server"

	uRequest "github.com/Dall06/go-cleanarch-template/src/pkg/user/delivery/handler"
	uMysql "github.com/Dall06/go-cleanarch-template/src/pkg/user/repository/mysql"
	uUCase "github.com/Dall06/go-cleanarch-template/src/pkg/user/usecase"
	"github.com/gorilla/mux"
)

func StartRouter(router *mux.Router, conn *sql.DB) {
	userMysqlRepo := uMysql.MySQLUserRepository(conn)
	userInteractor := uUCase.NewUserInteractor(userMysqlRepo)
	userHandler := uRequest.NewUserHandler(userInteractor)

	amw := middleware.NewAuthenticationMiddleWre()
	amw.Populate()

	cors := middleware.NewCORSMiddleware("GET, POST, PUT, DELETE", "localhost")
	
	routes.NewUserRoutes(router, userHandler).SetUserRoutes()

	router.Use(amw.Middleware)
	router.Use(cors.EnableCORS)

	server.NewGracefullyShutDown(router, config.Port).RunGracefully()
}
