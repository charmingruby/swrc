package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_hasPermission() {
	githubDisplayName := "charmingruby"
	email := "dummy@email.com"
	password := "password123"

	s.Run("it should be able to validate a user with a role", func() {
		acc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
				IsValid:           true,
				Verified:          true,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		err = s.accountUseCase.hasPermission(acc.ID, "account", entity.ACCOUNT_ROLE_MANAGER)
		s.NoError(err)
	})

	s.Run("it should be able to validate a user without a role", func() {
		acc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				Role:              entity.ACCOUNT_ROLE_DEVELOPER,
				IsValid:           true,
				Verified:          true,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		err = s.accountUseCase.hasPermission(acc.ID, "account", entity.ACCOUNT_ROLE_MANAGER)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to validate a user that don't exists", func() {
		err := s.accountUseCase.hasPermission("invalid id", "account", entity.ACCOUNT_ROLE_MANAGER)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to validate a user if solicitor is not a valid account", func() {
		acc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
				IsValid:           false,
				Verified:          true,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		err = s.accountUseCase.hasPermission(acc.ID, "account", entity.ACCOUNT_ROLE_MANAGER)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})
}
