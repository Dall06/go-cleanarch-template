package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user"
	"gopkg.in/go-playground/validator.v9"
)

type UserHelper struct {
	Validator *validator.Validate
}

func NewUserHelper() *UserHelper {
	return &UserHelper{
		Validator: validator.New(),
	}
}

func (uh *UserHelper) ValidateIndexRequest(r *http.Request) (string, error) {
	var index IndexUser

	err := json.NewDecoder(r.Body).Decode(&index)
	if err != nil {
		return "", nil
	}

	email := index.Email

	return email, nil
}

func (uh *UserHelper) ValidateSaveRequest(r *http.Request) (internal_user.UserAccount, error) {
	var save Save
	var uAccount internal_user.UserAccount

	err := json.NewDecoder(r.Body).Decode(&save)
	if err != nil {
		return uAccount, err
	}

	err = uh.Validator.Struct(save)
	if err != nil {
		return uAccount, err
	}

	uAccount = internal_user.UserAccount{
		Email:    save.Email,
		Password: save.Password,
		Phone:    save.Phone,
		Data: internal_user.UserData{
			Name:     save.Name,
			LastName: save.LastName,
			Region:   save.DataRegion,
		},
		Plan: save.PlanID,
	}

	return uAccount, nil
}

func (uh *UserHelper) ValidateChangeRequest(r *http.Request) (internal_user.UserAccount, string, error) {
	var change Change
	var uAccount internal_user.UserAccount
	var newEmail string

	err := json.NewDecoder(r.Body).Decode(&change)
	if err != nil {
		return uAccount, newEmail, err
	}

	err = uh.Validator.Struct(change)
	if err != nil {
		return uAccount, newEmail, err
	}

	uAccount = internal_user.UserAccount{
		Email: change.Email,
		Phone: change.Phone,
		Data: internal_user.UserData{
			Name:     change.Name,
			LastName: change.LastName,
			Region:   change.DataRegion,
		},
	}
	newEmail = change.NewEmail

	return uAccount, newEmail, err
}

func (uh *UserHelper) ValidateChangePlanRequest(r *http.Request) (internal_user.UserAccount, error) {
	var change ChangePlan
	var uAccount internal_user.UserAccount

	err := json.NewDecoder(r.Body).Decode(&change)
	if err != nil {
		return uAccount, err
	}

	err = uh.Validator.Struct(change)
	if err != nil {
		return uAccount, err
	}

	uAccount = internal_user.UserAccount{
		Email: change.Email,
		Plan: change.PlanID,
	}

	return uAccount, err
}

func (uh *UserHelper) ValidateChangePasswordRequest(r *http.Request) (internal_user.UserAccount, string, error) {
	var change ChangePass
	var uAccount internal_user.UserAccount
	var pass string

	err := json.NewDecoder(r.Body).Decode(&change)
	if err != nil {
		return uAccount, pass, err
	}

	err = uh.Validator.Struct(change)
	if err != nil {
		return uAccount, pass, err
	}

	uAccount = internal_user.UserAccount{
		Email: change.Email,
		Password: change.Password,
	}
	pass = change.NewPassword

	return uAccount, pass, err
}

func (uh *UserHelper) ValidateDestroyRequest(r *http.Request) (internal_user.UserAccount, error) {
	var destroy Destroy
	var uAccount internal_user.UserAccount

	err := json.NewDecoder(r.Body).Decode(&destroy)
	if err != nil {
		return uAccount, err
	}

	err = uh.Validator.Struct(destroy)
	if err != nil {
		return uAccount, err
	}

	uAccount = internal_user.UserAccount{
		Email: destroy.Email,
		Password: destroy.Password,
	}

	return uAccount, err
}