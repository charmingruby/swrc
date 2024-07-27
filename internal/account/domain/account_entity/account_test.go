package account_entity

import (
	"testing"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewAccount(t *testing.T) {
	ghDisplayName := "charmingruby"
	email := "dummy@mail.com"
	password := "password123"

	t.Run("it should be able to create a new account", func(t *testing.T) {
		account, err := NewAccount(ghDisplayName, email, password)

		assert.NoError(t, err)
		assert.Equal(t, ghDisplayName, account.GithubDisplayName)
		assert.Equal(t, email, account.Email)
		assert.Equal(t, false, account.Verified)
		assert.Equal(t, ACCOUNT_ROLE_DEVELOPER, account.Role)
		assert.Equal(t, password, account.Password)
	})

	t.Run("it should be not able to create an account with invalid email format", func(t *testing.T) {
		account, err := NewAccount(ghDisplayName, "invalid email", password)

		assert.Error(t, err)
		assert.Nil(t, account)
		assert.Equal(t, core.NewValidationErr(core.ErrInvalidFormat("email")).Error(), err.Error())
	})

	t.Run("it should be not able to create an account with blank github display name", func(t *testing.T) {
		account, err := NewAccount("", email, password)

		assert.Error(t, err)
		assert.Nil(t, account)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("githubdisplayname")).Error(), err.Error())
	})

	t.Run("it should be not able to create an account with less than minimum password required characters", func(t *testing.T) {
		account, err := NewAccount(ghDisplayName, email, "1234")

		assert.Error(t, err)
		assert.Nil(t, account)
		assert.Equal(t, core.NewValidationErr(core.ErrMinLength("password", "8")).Error(), err.Error())
	})

	t.Run("it should be not able to create an account with greater than maximum password required characters", func(t *testing.T) {
		account, err := NewAccount(ghDisplayName, email, "12345678123456781")

		assert.Error(t, err)
		assert.Nil(t, account)
		assert.Equal(t, core.NewValidationErr(core.ErrMaxLength("password", "16")).Error(), err.Error())
	})
}
