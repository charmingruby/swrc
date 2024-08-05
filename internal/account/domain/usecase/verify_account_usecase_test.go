package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_VerifyAccountUseCase() {
	accountToVerifyDisplayName := "charmingruby"
	accountToVerifyEmail := "dummy@email.com"
	accountToVerifyPassword := "password"

	solicitorGhDisplayName := accountToVerifyDisplayName + "-solic"
	solicitorEmail := accountToVerifyEmail + "-solic"
	solicitorPassword := accountToVerifyPassword + "-solic"

	s.Run("it should be able to verify an account", func() {
		isVerified := true
		isValid := true

		solicitorAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: solicitorGhDisplayName,
				Email:             solicitorEmail,
				Password:          solicitorPassword,
				IsValid:           &isValid,
				Verified:          &isVerified,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
			},
		)
		s.NoError(err)

		accToVerify, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: accountToVerifyDisplayName,
				Email:             accountToVerifyEmail,
				Password:          accountToVerifyPassword,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
			},
		)
		s.NoError(err)

		verification := true
		input := dto.VerifyAccountInputDTO{
			SolicitorAccountID: solicitorAcc.ID,
			AccountToVerifyID:  accToVerify.ID,
			Verification:       verification,
		}

		err = s.accountUseCase.VerifyAccountUseCase(input)
		s.NoError(err)

		newVerifiedAcc, err := s.accountRepository.FindByID(accToVerify.ID)
		s.NoError(err)
		s.Equal(verification, newVerifiedAcc.Verification.IsValid)
		s.True(newVerifiedAcc.Verification.Verified)
	})

	s.Run("it should be able to verify an account even if already have verification values", func() {
		isVerified := true
		isValid := true

		solicitorAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: solicitorGhDisplayName,
				Email:             solicitorEmail,
				Password:          solicitorPassword,
				IsValid:           &isValid,
				Verified:          &isVerified,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
			},
		)
		s.NoError(err)

		accToVerify, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: accountToVerifyDisplayName,
				Email:             accountToVerifyEmail,
				Password:          accountToVerifyPassword,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
				IsValid:           &isValid,
				Verified:          &isVerified,
			},
		)
		s.NoError(err)

		verification := false
		input := dto.VerifyAccountInputDTO{
			SolicitorAccountID: solicitorAcc.ID,
			AccountToVerifyID:  accToVerify.ID,
			Verification:       verification,
		}

		err = s.accountUseCase.VerifyAccountUseCase(input)
		s.NoError(err)

		newVerifiedAcc, err := s.accountRepository.FindByID(accToVerify.ID)
		s.NoError(err)
		s.Equal(verification, newVerifiedAcc.Verification.IsValid)
		s.True(newVerifiedAcc.Verification.Verified)
	})

	s.Run("it should be not able to verify an account when solicitor account id is invalid", func() {
		accToVerify, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: accountToVerifyDisplayName,
				Email:             accountToVerifyEmail,
				Password:          accountToVerifyPassword,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
			},
		)
		s.NoError(err)

		verification := true
		input := dto.VerifyAccountInputDTO{
			SolicitorAccountID: "invalid id",
			AccountToVerifyID:  accToVerify.ID,
			Verification:       verification,
		}

		err = s.accountUseCase.VerifyAccountUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("solicitor account").Error(), err.Error())
	})

	s.Run("it should be not able to verify an account when solicitor don't have needed permissions", func() {
		isVerified := true
		isValid := true

		solicitorAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: solicitorGhDisplayName,
				Email:             solicitorEmail,
				Password:          solicitorPassword,
				IsValid:           &isValid,
				Verified:          &isVerified,
				Role:              entity.ACCOUNT_ROLE_DEVELOPER,
			},
		)
		s.NoError(err)

		accToVerify, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: accountToVerifyDisplayName,
				Email:             accountToVerifyEmail,
				Password:          accountToVerifyPassword,
				Role:              entity.ACCOUNT_ROLE_DEVELOPER,
			},
		)
		s.NoError(err)

		verification := true
		input := dto.VerifyAccountInputDTO{
			SolicitorAccountID: solicitorAcc.ID,
			AccountToVerifyID:  accToVerify.ID,
			Verification:       verification,
		}

		err = s.accountUseCase.VerifyAccountUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to verify an account when account to verify id is invalid", func() {
		isVerified := true
		isValid := true

		solicitorAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: solicitorGhDisplayName,
				Email:             solicitorEmail,
				Password:          solicitorPassword,
				IsValid:           &isValid,
				Verified:          &isVerified,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
			},
		)
		s.NoError(err)

		verification := true
		input := dto.VerifyAccountInputDTO{
			SolicitorAccountID: solicitorAcc.ID,
			AccountToVerifyID:  "invalid id",
			Verification:       verification,
		}

		err = s.accountUseCase.VerifyAccountUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account to verify").Error(), err.Error())
	})

	s.Run("it should be not able to verify an account when account to verify is has already that verification value", func() {
		isVerified := true
		isValid := true

		solicitorAcc, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: solicitorGhDisplayName,
				Email:             solicitorEmail,
				Password:          solicitorPassword,
				IsValid:           &isValid,
				Verified:          &isVerified,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
			},
		)
		s.NoError(err)

		accToVerify, err := factory.MakeAccount(
			s.accountRepository,
			factory.MakeAccountInput{
				GithubDisplayName: accountToVerifyDisplayName,
				Email:             accountToVerifyEmail,
				Password:          accountToVerifyPassword,
				Role:              entity.ACCOUNT_ROLE_MANAGER,
				IsValid:           &isValid,
				Verified:          &isVerified,
			},
		)
		s.NoError(err)

		verification := true
		input := dto.VerifyAccountInputDTO{
			SolicitorAccountID: solicitorAcc.ID,
			AccountToVerifyID:  accToVerify.ID,
			Verification:       verification,
		}

		err = s.accountUseCase.VerifyAccountUseCase(input)
		s.Error(err)
		s.Equal(core.NewNothingToChangeErr().Error(), err.Error())
	})
}
