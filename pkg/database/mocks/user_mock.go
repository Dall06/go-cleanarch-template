package mocks

import (
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dall06/go-cleanarch-template/pkg/database"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user"
	"github.com/Dall06/go-cleanarch-template/pkg/internal/internal_user/repository"
)

type userhMocks struct {
	mock *database.Mock
}

func NewUserMocks(m *database.Mock) *userhMocks {
	return &userhMocks{
		mock: m,
	}
}

func (um *userhMocks) SelectUserMock(u *internal_user.UserAccount) *sql.DB {
	rows := sqlmock.NewRows([]string{
		"u_email",
		"u_phone",
		"u_name",
		"u_lastname",
		"u_region",
	}).AddRow(
		&u.Email,
		&u.Phone,
		&u.Data.Name,
		&u.Data.LastName,
		&u.Data.Region,
	)

	um.mock.Sqlmock.ExpectQuery(regexp.QuoteMeta(repository.SelectUser)).
		WithArgs(&u.Email).
		WillReturnRows(rows)

	return um.mock.DB
}

func (um *userhMocks) SelectUserAndPlanMock(u *internal_user.UserAccount) *sql.DB {
	rows := sqlmock.NewRows([]string{
		"u_email",
		"u_phone",
		"u_name",
		"u_lastname",
		"u_region",
		"plan",
	}).AddRow(
		&u.Email,
		&u.Phone,
		&u.Data.Name,
		&u.Data.LastName,
		&u.Data.Region,
		&u.Plan,
	)

	um.mock.Sqlmock.ExpectQuery(regexp.QuoteMeta(repository.SelectUserAndPlan)).WithArgs(&u.Email).WillReturnRows(rows)

	return um.mock.DB
}

func (um *userhMocks) AddUserMock(u *internal_user.UserAccount) (*sql.DB, sqlmock.Sqlmock) {
	//mock.ExpectBegin()
	um.mock.Sqlmock.ExpectExec(regexp.QuoteMeta(repository.AddUser)).WithArgs(
		&u.Email,
		&u.Password,
		&u.Phone,
		&u.Data.Name,
		&u.Data.LastName,
		&u.Data.Region,
		&u.Plan,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.DB, um.mock.Sqlmock
}

func (um *userhMocks) UpdateUserMock(u *internal_user.UserAccount, email string) (*sql.DB, sqlmock.Sqlmock) {
	um.mock.Sqlmock.ExpectExec(regexp.QuoteMeta(repository.UpdateUser)).WithArgs(
		&u.Email,
		email,
		&u.Phone,
		&u.Data.Name,
		&u.Data.LastName,
		&u.Data.Region,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.DB, um.mock.Sqlmock
}

func (um *userhMocks) UpdatePlanMock(u *internal_user.UserAccount) (*sql.DB, sqlmock.Sqlmock) {
	//mock.ExpectBegin()
	um.mock.Sqlmock.ExpectExec(regexp.QuoteMeta(repository.UpdateUserPlan)).WithArgs(
		&u.Email,
		&u.Data.ID,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.DB, um.mock.Sqlmock
}

func (um *userhMocks) UpdatePasswordMock(u *internal_user.UserAccount, p string) (*sql.DB, sqlmock.Sqlmock) {
	//mock.ExpectBegin()
	um.mock.Sqlmock.ExpectExec(regexp.QuoteMeta(repository.UpdateUserPassword)).WithArgs(
		&u.Email,
		&u.Password,
		p,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.DB, um.mock.Sqlmock
}

func (um *userhMocks) DeleteUserMock(u *internal_user.UserAccount) (*sql.DB, sqlmock.Sqlmock) {
	//mock.ExpectBegin()
	um.mock.Sqlmock.ExpectExec(regexp.QuoteMeta(repository.DeleteUser)).WithArgs(
		&u.Email,
		&u.Password,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.DB, um.mock.Sqlmock
}
