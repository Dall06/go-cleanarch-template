// Added methos of Auth repo because they are related to user domain entity
package repo_test

import (
	"fmt"
	"testing"

	"github.com/Dall06/go-cleanarch-template/src/domain"
	"github.com/Dall06/go-cleanarch-template/src/pkg/user/repository/mysql"
	"github.com/Dall06/go-cleanarch-template/test/mocks"
)

var u = domain.UserAccount{
	Email: "User",
	Phone: "473TEST123",
	Data: domain.UserData{
		Name:     "User",
		LastName: "Test",
		Region:   "MX",
	},
	Plan: 1,
}


func TestSelectUser(t *testing.T) {
	mock := mocks.NewMock()
	um := mocks.NewUserMocks(mock)
	db := um.SelectUserMock(&u)

	userRepo := mysql.MySQLUserRepository(db)
	user, err := userRepo.SelectUser(u.Email)

	fmt.Println("ERROR ", err)
	fmt.Println("USER ", user)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSelectUserAndPlan(t *testing.T) {
	mock := mocks.NewMock()
	um := mocks.NewUserMocks(mock)
	db := um.SelectUserMock(&u)

	userRepo := mysql.MySQLUserRepository(db)
	user, err := userRepo.SelectUser(u.Email)

	fmt.Println("ERROR ", err)
	fmt.Println("USER ", user)

	if err != nil {
		t.Fatal(err)
	}
}

func TestAddUser(t *testing.T) {
	mock := mocks.NewMock()
	um := mocks.NewUserMocks(mock)
	db, _ := um.AddUserMock(&u)

	userRepo := mysql.MySQLUserRepository(db)
	user, err := userRepo.AddUser(&u)

	fmt.Println("ERROR ", err)
	fmt.Println("USER ", user)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateUser(t *testing.T) {
	email := "test2@gmail.com"
	mock := mocks.NewMock()
	um := mocks.NewUserMocks(mock)
	db, _ := um.UpdateUserMock(&u, email)

	userRepo := mysql.MySQLUserRepository(db)
	user, err := userRepo.UpdateUser(&u, email)

	fmt.Println("ERROR ", err)
	fmt.Println("USER ", user)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdatePlan(t *testing.T) {
	mock := mocks.NewMock()
	um := mocks.NewUserMocks(mock)
	db, _ := um.UpdatePlanMock(&u)

	userRepo := mysql.MySQLUserRepository(db)
	user, err := userRepo.NewPlan(&u)

	fmt.Println("ERROR ", err)
	fmt.Println("USER ", user)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdatePassword(t *testing.T) {
	pass := "Test4321"
	mock := mocks.NewMock()
	um := mocks.NewUserMocks(mock)
	db, _ := um.UpdatePasswordMock(&u, pass)

	userRepo := mysql.MySQLUserRepository(db)
	user, err := userRepo.NewPassword(&u, pass)

	fmt.Println("ERROR ", err)
	fmt.Println("USER ", user)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUser(t *testing.T) {
	mock := mocks.NewMock()
	um := mocks.NewUserMocks(mock)
	db, _ := um.DeleteUserMock(&u)

	userRepo := mysql.MySQLUserRepository(db)
	user, err := userRepo.DeleteUser(&u)

	fmt.Println("ERROR ", err)
	fmt.Println("USER ", user)

	if err != nil {
		t.Fatal(err)
	}
}
