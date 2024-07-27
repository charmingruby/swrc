package account_repository

import "github.com/charmingruby/swrc/internal/account/domain/account_entity"

type AccountRepository interface {
	FindByID(id string) (*account_entity.Account, error)
	FindByEmail(email string) (*account_entity.Account, error)
	FindByGithubDisplayName(githubDisplayName string) (*account_entity.Account, error)
	Store(acc *account_entity.Account) (string, error)
	SaveVerification(acc *account_entity.Account) error
	SaveRole(acc *account_entity.Account) error
}
