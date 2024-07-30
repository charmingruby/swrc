package account_usecase

import "github.com/charmingruby/swrc/internal/common/core"

func (s *AccountUseCaseRegistry) hasRole(accountID string, role string) (bool, error) {
	acc, err := s.AccountRepository.FindByID(accountID)
	if err != nil {
		return false, core.NewNotFoundErr("account")
	}

	if acc.Role != role {
		return false, nil
	}

	return true, nil
}
