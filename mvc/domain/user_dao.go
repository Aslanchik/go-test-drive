package domain

import (
	"fmt"
	"go-test-drive/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Koko", LastName: "Akof", Email: "koko@akof.com"},
	}
)

func GetUser(userId int64) (*User, *utils.AppplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.AppplicationError{
		Message:    fmt.Errorf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
