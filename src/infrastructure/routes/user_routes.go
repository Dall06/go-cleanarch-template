package routes

import (
	"net/http"

	"github.com/Dall06/go-cleanarch-template/config"
	"github.com/Dall06/go-cleanarch-template/src/pkg/user/delivery/handler"
	"github.com/gorilla/mux"
)

type userRoutes struct {
	Router *mux.Router
	UserHandler *handler.UserHandler
}

func NewUserRoutes(r *mux.Router, uh *handler.UserHandler) *userRoutes{
	return &userRoutes{
		Router: r,
		UserHandler: uh,
	}
}

func (ur *userRoutes) SetUserRoutes() {
	var pathPrefix_v1 string = config.RouterBasePath_V1 + "/user"
	subRoute := ur.Router.PathPrefix(pathPrefix_v1).Subrouter()
	subRoute.HandleFunc("/welcome", Welcome).Methods(http.MethodGet)
	subRoute.HandleFunc("/", ur.UserHandler.IndexUser).Methods(http.MethodGet)
	subRoute.HandleFunc("/plan", ur.UserHandler.IndexUserAndPlan).Methods(http.MethodGet)
	subRoute.HandleFunc("/plan", ur.UserHandler.ChangePlan).Methods(http.MethodPut)
	subRoute.HandleFunc("/save", ur.UserHandler.Save).Methods(http.MethodPost)
	subRoute.HandleFunc("/change", ur.UserHandler.Change).Methods(http.MethodPost)
	subRoute.HandleFunc("/change/P", ur.UserHandler.ChangePassword).Methods(http.MethodPut)
	subRoute.HandleFunc("/byebye", ur.UserHandler.Destroy).Methods(http.MethodDelete)
}