package account_usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/account_dto"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_AuthenticateUseCase() {
	githubDisplayName := "charmingruby"
	email := "dummy@email.com"
	password := "password123"

	s.Run("it should be able to authenticate", func() {
		acc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		input := account_dto.AuthenticateInputDTO{
			Email:    email,
			Password: password,
		}

		output, err := s.accountUseCase.AuthenticateUseCase(input)

		s.NoError(err)
		s.Equal(acc.ID, output.ID)
		s.Equal(acc.Verified, output.Verified)
		s.Equal(acc.Role, output.Role)
	})

	s.Run("it should be not able to authenticate with an invalid email", func() {
		input := account_dto.AuthenticateInputDTO{
			Email:    email,
			Password: password,
		}

		output, err := s.accountUseCase.AuthenticateUseCase(input)

		s.Error(err)
		s.Nil(output)
		s.Equal(core.NewInvalidCredentialsErr().Error(), err.Error())
	})

	s.Run("it should be not able to authenticate with an invalid password", func() {
		_, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		input := account_dto.AuthenticateInputDTO{
			Email:    email,
			Password: password + "1",
		}

		output, err := s.accountUseCase.AuthenticateUseCase(input)

		s.Error(err)
		s.Nil(output)
		s.Equal(core.NewInvalidCredentialsErr().Error(), err.Error())
	})
}
