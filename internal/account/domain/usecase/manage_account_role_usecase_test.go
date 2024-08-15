package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_ManageAccountRoleUseCase() {
	githubDisplayName := "charmingruby"
	email := "dummy@email.com"
	password := "password123"
	baseRole := entity.ACCOUNT_ROLE_DEVELOPER
	managerRole := entity.ACCOUNT_ROLE_MANAGER

	s.Run("it should be able to change an account role", func() {
		managerAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              managerRole,
			},
		)
		s.NoError(err)

		accToManage, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              baseRole,
			},
		)
		s.NoError(err)

		s.Equal(2, len(s.accountRepository.Items))

		newRole := entity.ACCOUNT_ROLE_MANAGER
		input := dto.ManageAccountRoleInputDTO{
			ManagerAccountID: managerAcc.ID,
			AccountID:        accToManage.ID,
			NewRole:          newRole,
		}

		err = s.useCase.ManageAccountRoleUseCase(input)
		s.NoError(err)

		modifiedAcc, err := s.accountRepository.FindByID(accToManage.ID)
		s.NoError(err)
		s.Equal(newRole, modifiedAcc.Role)
	})

	s.Run("it should be not able to change an account role if actor id is invalid", func() {
		accToManage, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              baseRole,
			},
		)
		s.NoError(err)

		s.Equal(1, len(s.accountRepository.Items))

		newRole := entity.ACCOUNT_ROLE_MANAGER
		input := dto.ManageAccountRoleInputDTO{
			ManagerAccountID: "invalid manager id",
			AccountID:        accToManage.ID,
			NewRole:          newRole,
		}

		err = s.useCase.ManageAccountRoleUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("manager account").Error(), err.Error())
	})

	s.Run("it should be not able to change an account role if actor is not a valid account", func() {
		managerAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           false,
				Verified:          true,
				Role:              managerRole,
			},
		)
		s.NoError(err)

		accToManage, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              baseRole,
			},
		)
		s.NoError(err)

		s.Equal(2, len(s.accountRepository.Items))

		newRole := entity.ACCOUNT_ROLE_MANAGER
		input := dto.ManageAccountRoleInputDTO{
			ManagerAccountID: managerAcc.ID,
			AccountID:        accToManage.ID,
			NewRole:          newRole,
		}

		err = s.useCase.ManageAccountRoleUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to change an account role if actor don't have the permissions", func() {
		managerAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              baseRole,
			},
		)
		s.NoError(err)

		accToManage, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              baseRole,
			},
		)
		s.NoError(err)

		s.Equal(2, len(s.accountRepository.Items))

		newRole := entity.ACCOUNT_ROLE_MANAGER
		input := dto.ManageAccountRoleInputDTO{
			ManagerAccountID: managerAcc.ID,
			AccountID:        accToManage.ID,
			NewRole:          newRole,
		}

		err = s.useCase.ManageAccountRoleUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to change an account role if account id is invalid", func() {
		managerAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              managerRole,
			},
		)
		s.NoError(err)

		s.Equal(1, len(s.accountRepository.Items))

		newRole := entity.ACCOUNT_ROLE_MANAGER
		input := dto.ManageAccountRoleInputDTO{
			ManagerAccountID: managerAcc.ID,
			AccountID:        "invalid account id",
			NewRole:          newRole,
		}

		err = s.useCase.ManageAccountRoleUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to change an account role if actor is not a valid account", func() {
		managerAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           false,
				Verified:          true,
				Role:              managerRole,
			},
		)
		s.NoError(err)

		accToManage, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              baseRole,
			},
		)
		s.NoError(err)

		s.Equal(2, len(s.accountRepository.Items))

		newRole := entity.ACCOUNT_ROLE_MANAGER
		input := dto.ManageAccountRoleInputDTO{
			ManagerAccountID: managerAcc.ID,
			AccountID:        accToManage.ID,
			NewRole:          newRole,
		}

		err = s.useCase.ManageAccountRoleUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to change an account role if new role is invalid", func() {
		managerAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              managerRole,
			},
		)
		s.NoError(err)

		accToManage, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              baseRole,
			},
		)
		s.NoError(err)

		s.Equal(2, len(s.accountRepository.Items))

		input := dto.ManageAccountRoleInputDTO{
			ManagerAccountID: managerAcc.ID,
			AccountID:        accToManage.ID,
			NewRole:          "invalid role",
		}

		err = s.useCase.ManageAccountRoleUseCase(input)
		s.Error(err)
		s.Equal(core.NewValidationErr("invalid role").Error(), err.Error())
	})

	s.Run("it should be not able to change an account role if new role is equal to the current role", func() {
		managerAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              managerRole,
			},
		)
		s.NoError(err)

		accToManage, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: githubDisplayName,
				Email:             email,
				Password:          password,
				IsValid:           true,
				Verified:          true,
				Role:              baseRole,
			},
		)
		s.NoError(err)

		s.Equal(2, len(s.accountRepository.Items))

		input := dto.ManageAccountRoleInputDTO{
			ManagerAccountID: managerAcc.ID,
			AccountID:        accToManage.ID,
			NewRole:          baseRole,
		}

		err = s.useCase.ManageAccountRoleUseCase(input)
		s.Error(err)
		s.Equal(core.NewNothingToChangeErr().Error(), err.Error())
	})
}
