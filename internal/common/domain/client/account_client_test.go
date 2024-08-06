package client

import (
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_AccountClient_AccountExists() {
	s.Run("it should get true if found an account even if is not validated", func() {
		acc, err := factory.MakeAccount(s.accountRepository, factory.MakeAccountInput{
			GithubDisplayName: "charmingruby",
			Email:             "dummy@email.com",
			Password:          "password123",
		})
		s.NoError(err)

		accExists := s.accountClient.AccountExists(acc.ID)

		s.True(accExists)
	})

	s.Run("it should get false if don't found an account", func() {
		accExists := s.accountClient.AccountExists("invalid id")

		s.False(accExists)
	})
}

func (s *Suite) Test_AccountClient_ValidAccountExists() {
	s.Run("it should receive no error if found an account", func() {
		isValid := true
		verified := true

		acc, err := factory.MakeAccount(s.accountRepository, factory.MakeAccountInput{
			GithubDisplayName: "charmingruby",
			Email:             "dummy@email.com",
			Password:          "password123",
			IsValid:           &isValid,
			Verified:          &verified,
		})
		s.NoError(err)

		err = s.accountClient.ValidAccountExists(acc.ID)
		s.NoError(err)
	})

	s.Run("it should receive an error if account doesn't exists", func() {
		err := s.accountClient.ValidAccountExists("invalid id")
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should receive an error if account is not valid", func() {
		isValid := false
		verified := false

		acc, err := factory.MakeAccount(s.accountRepository, factory.MakeAccountInput{
			GithubDisplayName: "charmingruby",
			Email:             "dummy@email.com",
			Password:          "password123",
			IsValid:           &isValid,
			Verified:          &verified,
		})
		s.NoError(err)

		err = s.accountClient.ValidAccountExists(acc.ID)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})
}

func (s *Suite) Test_AccountClient_ValidAccountExistsAndMatchRole() {
	s.Run("it should receive no error if found an account and match role", func() {
		isValid := true
		verified := true

		acc, err := factory.MakeAccount(s.accountRepository, factory.MakeAccountInput{
			GithubDisplayName: "charmingruby",
			Email:             "dummy@email.com",
			Password:          "password123",
			Role:              entity.ACCOUNT_ROLE_MANAGER,
			IsValid:           &isValid,
			Verified:          &verified,
		})
		s.NoError(err)

		err = s.accountClient.ValidAccountExistsAndMatchRole(acc.ID, entity.ACCOUNT_ROLE_MANAGER)
		s.NoError(err)
	})
	s.Run("it should receive an error if account doesn't exists", func() {
		err := s.accountClient.ValidAccountExistsAndMatchRole("invalid id", entity.ACCOUNT_ROLE_MANAGER)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should receive an error if account is invalid", func() {
		isValid := false
		verified := false

		acc, err := factory.MakeAccount(s.accountRepository, factory.MakeAccountInput{
			GithubDisplayName: "charmingruby",
			Email:             "dummy@email.com",
			Password:          "password123",
			Role:              entity.ACCOUNT_ROLE_MANAGER,
			IsValid:           &isValid,
			Verified:          &verified,
		})
		s.NoError(err)

		err = s.accountClient.ValidAccountExistsAndMatchRole(acc.ID, entity.ACCOUNT_ROLE_MANAGER)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should receive an error if account role doesn't match", func() {
		isValid := true
		verified := true

		acc, err := factory.MakeAccount(s.accountRepository, factory.MakeAccountInput{
			GithubDisplayName: "charmingruby",
			Email:             "dummy@email.com",
			Password:          "password123",
			Role:              entity.ACCOUNT_ROLE_DEVELOPER,
			IsValid:           &isValid,
			Verified:          &verified,
		})
		s.NoError(err)

		err = s.accountClient.ValidAccountExistsAndMatchRole(acc.ID, entity.ACCOUNT_ROLE_MANAGER)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})
}
