package routes

import (
	"net/http"

	"github.com/Dall06/go-cleanarch-template/src/infrastructure/server"
)

func Welcome(rw http.ResponseWriter, r *http.Request) {
	server.NewResponseHandler().RespondWithSuccess("welcome", rw)
}