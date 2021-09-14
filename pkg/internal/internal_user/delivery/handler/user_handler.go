package handler

import (
	"net/http"

	"github.com/Dall06/go-cleanarch-template/pkg/infrastructure/services"
	"github.com/Dall06/go-cleanarch-template/pkg/infrastructure/middleware"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user/delivery"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user/usecase"
)

type UserHandler struct {
	UserInteractor *usecase.UserInteractor
	ResponseHandler	*services.ResponseHandler
	JwtHandler	*middleware.JWTHandler
	LoggerHandler *services.LoggerHandler
	UserHelper	*delivery.UserHelper
}

func NewUserHandler(ai *usecase.UserInteractor) *UserHandler {
	return &UserHandler{
		UserInteractor: ai,
		ResponseHandler: services.NewResponseHandler(),
		JwtHandler: middleware.NewJWTHandler(),
		LoggerHandler: services.NewLoggerHandler(),
		UserHelper: delivery.NewUserHelper(),
	}
}

func (userHandler *UserHandler) IndexUser(w http.ResponseWriter, r *http.Request) {	
	email, err := userHandler.UserHelper.ValidateIndexRequest(r)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithBadRequest(err, w)
	}
	user, err := userHandler.UserInteractor.IndexUser(email)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
	}
	userHandler.LoggerHandler.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	userHandler.ResponseHandler.RespondWithSuccess(user, w)
}

func (userHandler *UserHandler) IndexUserAndPlan(w http.ResponseWriter, r *http.Request) {
	email, err := userHandler.UserHelper.ValidateIndexRequest(r)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithBadRequest(err, w)
	}

	user, err := userHandler.UserInteractor.IndexUserAndPlan(email)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
	}
	userHandler.LoggerHandler.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	userHandler.ResponseHandler.RespondWithSuccess(user, w)
}

func (userHandler *UserHandler) Save(w http.ResponseWriter, r *http.Request) {
	userToSave, err := userHandler.UserHelper.ValidateSaveRequest(r)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithBadRequest(err, w)
	}

	response, err := userHandler.UserInteractor.Save(&userToSave)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
	}
	userHandler.LoggerHandler.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	userHandler.ResponseHandler.RespondWithSuccess(response, w)
}

func (userHandler *UserHandler) Change(w http.ResponseWriter, r *http.Request) {
	validated, err := userHandler.JwtHandler.ValidateTokenCookie(r)
	if err != nil {
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
		userHandler.LoggerHandler.LogError("%s", err)
	}
	if !validated {
		userHandler.ResponseHandler.RespondWithUnauthorized(err, w)
		userHandler.LoggerHandler.LogError("%s NO VALIDATED", err)
	}

	userToChange, newEmail, err := userHandler.UserHelper.ValidateChangeRequest(r)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithBadRequest(err, w)
	}

	response, err := userHandler.UserInteractor.Change(&userToChange, newEmail)
	if err != nil {
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
		userHandler.LoggerHandler.LogError("%s", err)
	}

	userHandler.LoggerHandler.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	userHandler.ResponseHandler.RespondWithSuccess(response, w)
}

func (userHandler *UserHandler) ChangePlan(w http.ResponseWriter, r *http.Request) {
	validated, err := userHandler.JwtHandler.ValidateTokenCookie(r)
	if err != nil {
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
		userHandler.LoggerHandler.LogError("%s", err)
	}
	if !validated {
		userHandler.ResponseHandler.RespondWithUnauthorized(err, w)
		userHandler.LoggerHandler.LogError("%s NO VALIDATED", err)
	}

	userToChange, err := userHandler.UserHelper.ValidateChangePlanRequest(r)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithBadRequest(err, w)
	}

	response, err := userHandler.UserInteractor.ChangePlan(&userToChange)
	if err != nil {
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
		userHandler.LoggerHandler.LogError("%s", err)
	}

	userHandler.LoggerHandler.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	userHandler.ResponseHandler.RespondWithSuccess(response, w)
}

func (userHandler *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	validated, err := userHandler.JwtHandler.ValidateTokenCookie(r)
	if err != nil {
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
		userHandler.LoggerHandler.LogError("%s", err)
	}
	if !validated {
		userHandler.ResponseHandler.RespondWithUnauthorized(err, w)
		userHandler.LoggerHandler.LogError("%s NO VALIDATED", err)
	}

	user, pass, err := userHandler.UserHelper.ValidateChangePasswordRequest(r)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithBadRequest(err, w)
	}

	response, err := userHandler.UserInteractor.ChangePassword(&user, pass)
	if err != nil {
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
		userHandler.LoggerHandler.LogError("%s", err)
	}

	userHandler.LoggerHandler.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	userHandler.ResponseHandler.RespondWithSuccess(response, w)
}

func (userHandler *UserHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	validated, err := userHandler.JwtHandler.ValidateTokenCookie(r)
	if err != nil {
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
		userHandler.LoggerHandler.LogError("%s", err)
	}
	if !validated {
		userHandler.ResponseHandler.RespondWithUnauthorized(err, w)
		userHandler.LoggerHandler.LogError("%s NO VALIDATED", err)
	}

	user, err := userHandler.UserHelper.ValidateDestroyRequest(r)
	if err != nil {
		userHandler.LoggerHandler.LogError("%s", err)
		userHandler.ResponseHandler.RespondWithBadRequest(err, w)
	}

	response, err := userHandler.UserInteractor.Destroy(&user)
	if err != nil {
		userHandler.ResponseHandler.RespondWithInternalServerError(err, w)
		userHandler.LoggerHandler.LogError("%s", err)
	}
	userHandler.LoggerHandler.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	userHandler.ResponseHandler.RespondWithSuccess(response, w)
}
