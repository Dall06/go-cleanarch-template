package usecase

import (
	"database/sql"

	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user"
)

// A UserRepository belong to the usecases layer.
type UserRepository interface {
	SelectUser(email string) (*internal_user.UserAccount, error)
	SelectUserAndPlan(string) (*internal_user.UserAccount, error)
	AddUser(*internal_user.UserAccount) (*sql.Result, error)
	UpdateUser(*internal_user.UserAccount, string) (*sql.Result, error)
	NewPlan(*internal_user.UserAccount) (*sql.Result, error)
	NewPassword(*internal_user.UserAccount, string) (*sql.Result, error)
	DeleteUser(*internal_user.UserAccount) (*sql.Result, error)
}
