package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser (t *testing.T) {
	user, err := NewUser("LucasCabral", "l@email.com", "Str0ngP4s5word")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "LucasCabral", user.Name)
	assert.Equal(t, "l@email.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("LucasCabral", "l@email.com", "Str0ngP4s5word")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("Str0ngP4s5word"))
	assert.False(t, user.ValidatePassword("for wrong passwords"))
	assert.NotEqual(t, "Str0ngP4s5word", user.Password)
}
