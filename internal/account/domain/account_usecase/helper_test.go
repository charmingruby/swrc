package account_usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/account_entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_hasRole() {
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
				Role:              account_entity.ACCOUNT_ROLE_MANAGER,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		hasPermissions, err := s.accountUseCase.hasRole(acc.ID, account_entity.ACCOUNT_ROLE_MANAGER)
		s.NoError(err)
		s.True(hasPermissions)
	})

	s.Run("it should be able to validate a user without a role", func() {
		acc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				Role:              account_entity.ACCOUNT_ROLE_DEVELOPER,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		hasPermissions, err := s.accountUseCase.hasRole(acc.ID, account_entity.ACCOUNT_ROLE_MANAGER)
		s.NoError(err)
		s.False(hasPermissions)
	})

	s.Run("it should be not able to validate a user that don't exists", func() {
		hasPermissions, err := s.accountUseCase.hasRole("invalid id", account_entity.ACCOUNT_ROLE_MANAGER)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
		s.False(hasPermissions)
	})
}
