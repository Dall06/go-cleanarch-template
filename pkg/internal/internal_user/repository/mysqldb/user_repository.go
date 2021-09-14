package mysqldb

import (
	"database/sql"

	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user/repository"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user/usecase"
)

// A UserRepository belong to the inteface layer
type mysqlUserRepository struct {
	DB *sql.DB
}

func MySQLUserRepository(db *sql.DB) usecase.UserRepository {
	return &mysqlUserRepository{DB: db}
}

func (ur *mysqlUserRepository) SelectUser(email string) (*internal_user.UserAccount, error) {
	var userAccount internal_user.UserAccount

	err := ur.DB.QueryRow(repository.SelectUser, email).Scan(
		&userAccount.Email,
		&userAccount.Phone,
		&userAccount.Data.Name,
		&userAccount.Data.LastName,
		&userAccount.Data.Region)

	if err != nil {
		return &userAccount, err
	}
	return &userAccount, nil
}

func (ur *mysqlUserRepository) SelectUserAndPlan(email string) (*internal_user.UserAccount, error) {
	var userAccount internal_user.UserAccount

	row := ur.DB.QueryRow(repository.SelectUserAndPlan, email)
	err := row.Scan(
		&userAccount.Email,
		&userAccount.Phone,
		&userAccount.Data.Name,
		&userAccount.Data.LastName,
		&userAccount.Data.Region,
		&userAccount.Plan)

	if err != nil {
		return &userAccount, err
	}

	return &userAccount, nil
}

func (ur *mysqlUserRepository) AddUser(userAccount *internal_user.UserAccount) (*sql.Result, error) {
	result, err := ur.DB.Exec(repository.AddUser,
		&userAccount.Email,
		&userAccount.Password,
		&userAccount.Phone,
		&userAccount.Data.Name,
		&userAccount.Data.LastName,
		&userAccount.Data.Region,
		&userAccount.Plan)

	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (ur *mysqlUserRepository) UpdateUser(userAccount *internal_user.UserAccount, newEmail string) (*sql.Result, error) {
	// 0 is New, 1 is old
	result, err := ur.DB.Exec(repository.UpdateUser,
		userAccount.Email,
		newEmail,
		userAccount.Phone,
		userAccount.Data.Name,
		userAccount.Data.LastName,
		userAccount.Data.Region)

	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (ur *mysqlUserRepository) NewPlan(userAccount *internal_user.UserAccount) (*sql.Result, error) {
	result, err := ur.DB.Exec(repository.UpdateUserPlan,
		&userAccount.Email,
		&userAccount.Plan)

	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (ur *mysqlUserRepository) NewPassword(ua *internal_user.UserAccount, newP string) (*sql.Result, error) {
	result, err := ur.DB.Exec(repository.UpdateUserPassword,
		&ua.Email,
		&ua.Password,
		&newP)

	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (ur *mysqlUserRepository) DeleteUser(userAccount *internal_user.UserAccount) (*sql.Result, error) {
	result, err := ur.DB.Exec(repository.DeleteUser,
		userAccount.Email,
		userAccount.Password)

	if err != nil {
		return &result, err
	}

	return &result, err
}
