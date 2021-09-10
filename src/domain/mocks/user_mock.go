package mocks

import (
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dall06/go-cleanarch-template/src/database/queries"
	"github.com/Dall06/go-cleanarch-template/src/domain"
)

type userhMocks struct {
	mock *Mock
}

func NewUserMocks(m *Mock) *userhMocks {
	return &userhMocks{
		mock: m,
	}
}

func (um *userhMocks) LogInMock(u *domain.UserAccount) *sql.DB {
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

	um.mock.mock.ExpectQuery(regexp.QuoteMeta(queries.LogIn)).
		WithArgs(u.Email, u.Password).
		WillReturnRows(rows)

	return um.mock.db
}

func (um *userhMocks) SelectUserMock(u *domain.UserAccount) *sql.DB {
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

	um.mock.mock.ExpectQuery(regexp.QuoteMeta(queries.SelectUser)).
		WithArgs(&u.Email).
		WillReturnRows(rows)

	return um.mock.db
}

func (um *userhMocks) SelectUserAndPlanMock(u *domain.UserAccount) *sql.DB {
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

	um.mock.mock.ExpectQuery(regexp.QuoteMeta(queries.SelectUserAndPlan)).WithArgs(&u.Email).WillReturnRows(rows)

	return um.mock.db
}

func (um *userhMocks) AddUserMock(u *domain.UserAccount) (*sql.DB, sqlmock.Sqlmock) {
	//mock.ExpectBegin()
	um.mock.mock.ExpectExec(regexp.QuoteMeta(queries.AddUser)).WithArgs(
		&u.Email,
		&u.Password,
		&u.Phone,
		&u.Data.Name,
		&u.Data.LastName,
		&u.Data.Region,
		&u.Plan,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.db, um.mock.mock
}

func (um *userhMocks) UpdateUserMock(u *domain.UserAccount, email string) (*sql.DB, sqlmock.Sqlmock) {
	um.mock.mock.ExpectExec(regexp.QuoteMeta(queries.UpdateUser)).WithArgs(
		&u.Email,
		email,
		&u.Phone,
		&u.Data.Name,
		&u.Data.LastName,
		&u.Data.Region,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.db, um.mock.mock
}

func (um *userhMocks) UpdatePlanMock(u *domain.UserAccount) (*sql.DB, sqlmock.Sqlmock) {
	//mock.ExpectBegin()
	um.mock.mock.ExpectExec(regexp.QuoteMeta(queries.UpdateUserPlan)).WithArgs(
		&u.Email,
		&u.Data.ID,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.db, um.mock.mock
}

func (um *userhMocks) UpdatePasswordMock(u *domain.UserAccount, p string) (*sql.DB, sqlmock.Sqlmock) {
	//mock.ExpectBegin()
	um.mock.mock.ExpectExec(regexp.QuoteMeta(queries.UpdateUserPassword)).WithArgs(
		&u.Email,
		&u.Password,
		p,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.db, um.mock.mock
}

func (um *userhMocks) DeleteUserMock(u *domain.UserAccount) (*sql.DB, sqlmock.Sqlmock) {
	//mock.ExpectBegin()
	um.mock.mock.ExpectExec(regexp.QuoteMeta(queries.DeleteUser)).WithArgs(
		&u.Email,
		&u.Password,
	).WillReturnResult(sqlmock.NewResult(0, 0))
	//mock.ExpectCommit()

	return um.mock.db, um.mock.mock
}
