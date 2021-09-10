package mysqldb

import (
	"database/sql"

	"github.com/Dall06/go-cleanarch-template/src/database/queries"
	"github.com/Dall06/go-cleanarch-template/src/domain"
	"github.com/Dall06/go-cleanarch-template/src/pkg/user/usecase"
)

// A UserRepository belong to the inteface layer
type mysqlUserRepository struct {
	DB *sql.DB
}

func MySQLUserRepository(db *sql.DB) usecase.UserRepository {
	return &mysqlUserRepository{DB: db}
}

func (ur *mysqlUserRepository) SelectUser(email string) (*domain.UserAccount, error) {
	var userAccount domain.UserAccount

	err := ur.DB.QueryRow(queries.SelectUser, email).Scan(
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

func (ur *mysqlUserRepository) SelectUserAndPlan(email string) (*domain.UserAccount, error) {
	var userAccount domain.UserAccount

	row := ur.DB.QueryRow(queries.SelectUserAndPlan, email)
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

func (ur *mysqlUserRepository) AddUser(userAccount *domain.UserAccount) (*sql.Result, error) {
	result, err := ur.DB.Exec(queries.AddUser,
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

func (ur *mysqlUserRepository) UpdateUser(userAccount *domain.UserAccount, newEmail string) (*sql.Result, error) {
	// 0 is New, 1 is old
	result, err := ur.DB.Exec(queries.UpdateUser,
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

func (ur *mysqlUserRepository) NewPlan(userAccount *domain.UserAccount) (*sql.Result, error) {
	result, err := ur.DB.Exec(queries.UpdateUserPlan,
		&userAccount.Email,
		&userAccount.Plan)

	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (ur *mysqlUserRepository) NewPassword(ua *domain.UserAccount, newP string) (*sql.Result, error) {
	result, err := ur.DB.Exec(queries.UpdateUserPassword,
		&ua.Email,
		&ua.Password,
		&newP)

	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (ur *mysqlUserRepository) DeleteUser(userAccount *domain.UserAccount) (*sql.Result, error) {
	result, err := ur.DB.Exec(queries.DeleteUser,
		userAccount.Email,
		userAccount.Password)

	if err != nil {
		return &result, err
	}

	return &result, err
}
