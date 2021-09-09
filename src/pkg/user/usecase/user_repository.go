package usecase

import (
	"database/sql"

	"github.com/Dall06/go-cleanarch-template/src/domain"
)

// A UserRepository belong to the usecases layer.
type UserRepository interface {
	SelectUser(email string) (*domain.UserAccount, error)
	SelectUserAndPlan(string) (*domain.UserAccount, error)
	AddUser(*domain.UserAccount) (*sql.Result, error)
	UpdateUser(*domain.UserAccount, string) (*sql.Result, error)
	NewPlan(*domain.UserAccount) (*sql.Result, error)
	NewPassword(*domain.UserAccount, string) (*sql.Result, error)
	DeleteUser(*domain.UserAccount) (*sql.Result, error)
}
