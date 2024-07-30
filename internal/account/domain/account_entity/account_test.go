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
		assert.Equal(t, false, account.Verification.Verified)
		assert.Equal(t, false, account.Verification.IsValid)
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

func Test_AccountVerify(t *testing.T) {
	ghDisplayName := "charmingruby"
	email := "dummy@mail.com"
	password := "password123"

	t.Run("it should be able to verify for the first time", func(t *testing.T) {
		acc, err := NewAccount(ghDisplayName, email, password)
		assert.NoError(t, err)

		err = acc.Verify(false)
		assert.NoError(t, err)
		assert.Equal(t, false, acc.Verification.IsValid)
		assert.Equal(t, true, acc.Verification.Verified)
	})

	t.Run("it should be able to modify verification even if it's already verification", func(t *testing.T) {
		acc, err := NewAccount(ghDisplayName, email, password)
		assert.NoError(t, err)

		err = acc.Verify(false)
		assert.NoError(t, err)
		assert.Equal(t, false, acc.Verification.IsValid)
		assert.Equal(t, true, acc.Verification.Verified)

		err = acc.Verify(true)
		assert.NoError(t, err)
		assert.Equal(t, true, acc.Verification.IsValid)
		assert.Equal(t, true, acc.Verification.Verified)
	})

	t.Run("it should be not able to verify with same verification", func(t *testing.T) {
		acc, err := NewAccount(ghDisplayName, email, password)
		assert.NoError(t, err)

		err = acc.Verify(false)
		assert.NoError(t, err)
		assert.Equal(t, false, acc.Verification.IsValid)
		assert.Equal(t, true, acc.Verification.Verified)

		err = acc.Verify(false)
		assert.Error(t, err)
		assert.Equal(t, core.NewValidationErr("nothing to change").Error(), err.Error())
		assert.Equal(t, false, acc.Verification.IsValid)
		assert.Equal(t, true, acc.Verification.Verified)
	})
}
