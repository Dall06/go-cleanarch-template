package routes

import (
	"net/http"

	"github.com/Dall06/go-cleanarch-template/pkg/infrastructure/services"
)

func Welcome(rw http.ResponseWriter, r *http.Request) {
	services.NewResponseHandler().RespondWithSuccess("welcome", rw)
}