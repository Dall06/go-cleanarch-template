package usecase

import (
	"database/sql"

	"github.com/Dall06/go-cleanarch-template/src/domain"
)

// A UserInteractor belong to the usecases layer.
type UserInteractor struct {
	UserRepository UserRepository
}

func NewUserInteractor(ur UserRepository) *UserInteractor{
	return &UserInteractor{
		UserRepository: ur,
	}
}

func (userInteractor *UserInteractor) IndexUser(email string) (user *domain.UserAccount, err error) {
	user, err = userInteractor.UserRepository.SelectUser(email)
	return
}

func (userInteractor *UserInteractor) IndexUserAndPlan(email string) (user *domain.UserAccount, err error) {
	user, err = userInteractor.UserRepository.SelectUserAndPlan(email)
	return
}

func (ui *UserInteractor) Save(user *domain.UserAccount) (result *sql.Result, err error) {
	result, err = ui.UserRepository.AddUser(user)
	return
}

func (ui *UserInteractor) Change(user *domain.UserAccount, email string) (result *sql.Result, err error) {
	result, err = ui.UserRepository.UpdateUser(user, email)
	return
}

func (ui *UserInteractor) ChangePlan(userAccount *domain.UserAccount) (result *sql.Result, err error) {
	result, err = ui.UserRepository.NewPlan(userAccount)
	return
}

func (ui *UserInteractor) ChangePassword(ua *domain.UserAccount, newP string) (result *sql.Result, err error) {
	result, err = ui.UserRepository.NewPassword(ua, newP)
	return
}

func (ui *UserInteractor) Destroy(userAccount *domain.UserAccount) (result *sql.Result, err error) {
	result, err = ui.UserRepository.DeleteUser(userAccount)
	return
}
