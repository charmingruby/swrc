package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_RegisterUseCase() {
	githubDisplayName := "charmingruby"
	email := "dummy@email.com"
	password := "password123"

	s.Run("it should be able to create a new account", func() {
		input := dto.RegisterInputDTO{
			GithubDisplayName: githubDisplayName,
			Email:             email,
			Password:          password,
		}

		output, err := s.accountUseCase.RegisterUseCase(input)
		s.NoError(err)

		repoAcc, err := s.accountRepository.FindByID(output.ID)
		s.NoError(err)
		s.Equal(githubDisplayName, repoAcc.GithubDisplayName)
		s.Equal(email, repoAcc.Email)

		hashedPassword, err := s.fakeHashService.GenerateHash(password)
		s.NoError(err)
		s.Equal(hashedPassword, repoAcc.Password)
	})

	s.Run("it should be not able to create a account with a conflicting email", func() {
		_, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		input := dto.RegisterInputDTO{
			GithubDisplayName: githubDisplayName + "1",
			Email:             email,
			Password:          password,
		}

		output, err := s.accountUseCase.RegisterUseCase(input)

		s.Error(err)
		s.Nil(output)
		s.Equal(core.NewConflictErr("account", "email").Error(), err.Error())
	})

	s.Run("it should be not able to create an account with a conflicting github display name", func() {
		_, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
			})
		s.NoError(err)
		s.Equal(1, len(s.accountRepository.Items))

		input := dto.RegisterInputDTO{
			GithubDisplayName: githubDisplayName,
			Email:             email + ".br",
			Password:          password,
		}

		output, err := s.accountUseCase.RegisterUseCase(input)

		s.Error(err)
		s.Nil(output)
		s.Equal(core.NewConflictErr("account", "github_display_name").Error(), err.Error())
	})

	s.Run("it should be not able to create an account with entity errors", func() {
		input := dto.RegisterInputDTO{
			GithubDisplayName: githubDisplayName,
			Email:             "",
			Password:          password,
		}

		output, err := s.accountUseCase.RegisterUseCase(input)

		s.Error(err)
		s.Nil(output)
		s.Equal(core.NewValidationErr(core.ErrRequired("email")).Error(), err.Error())
	})
}
