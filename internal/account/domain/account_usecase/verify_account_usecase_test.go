package account_usecase

func (s *Suite) Test_VerifyAccountUseCase() {
	s.Run("it should be able to verify an account", func() {})

	s.Run("it should be not able to verify an account when solicitor account id is invalid", func() {
	})

	s.Run("it should be not able to verify an account when solicitor don't have needed permissions", func() {})

	s.Run("it should be not able to verify an account when account to verify id is invalid", func() {})

	s.Run("it should be not able to verify an account when account to verify id is already verified", func() {})
}
