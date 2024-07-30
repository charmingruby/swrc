package usecase

import "github.com/charmingruby/swrc/internal/common/core"

func (s *AccountUseCaseRegistry) hasPermission(accountID string, role string) (bool, error) {
	acc, err := s.AccountRepository.FindByID(accountID)
	if err != nil {
		return false, core.NewNotFoundErr("account")
	}

	if !acc.Verification.IsValid {
		return false, core.NewUnauthorizedErr()
	}

	if acc.Role != role {
		return false, nil
	}

	return true, nil
}
