package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := GetUser(0)

	// Validation:
	assert.Nil(t, user, "we are not expecting a user with id 0")
	assert.NotNil(t, err, "we are expecting an error when user id is 0")
	assert.EqualValues(t, err.Status, http.StatusNotFound, "we are expecting 404 when user is not found")
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)

	// Those are the same
	if user != nil {
		t.Error("we are not expecting a user with id 0")
	}
	if err == nil {
		t.Error("we are expecting an error when user id is 0")
	}
	if err.Status != http.StatusNotFound {
		t.Error("we are expecting 404 when user is not found")
	}
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Jorge", user.FirstName)
	assert.EqualValues(t, "González", user.LastName)
	assert.EqualValues(t, "jgilbaja@gmail.com", user.Email)
}
