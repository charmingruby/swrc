package usecase

import "github.com/charmingruby/swrc/internal/common/core"

func (s *AccountUseCaseRegistry) hasPermission(accountID, entityIdentifier, role string) error {
	acc, err := s.AccountRepository.FindByID(accountID)
	if err != nil {
		return core.NewNotFoundErr(entityIdentifier)
	}

	if !acc.Verification.IsValid {
		return core.NewUnauthorizedErr()
	}

	if acc.Role != role {
		return core.NewUnauthorizedErr()
	}

	return nil
}
