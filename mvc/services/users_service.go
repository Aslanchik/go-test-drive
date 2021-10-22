package services

import (
	"go-test-drive/mvc/domain"
	"go-test-drive/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.AppplicationError) {
	return domain.GetUser(userId)
}
