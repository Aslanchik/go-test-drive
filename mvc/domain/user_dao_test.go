package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Koko", user.FirstName)
	assert.EqualValues(t, "Akof", user.LastName)
	assert.EqualValues(t, "koko@akof.com", user.Email)
}
