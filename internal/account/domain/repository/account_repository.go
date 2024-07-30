package repository

import "github.com/charmingruby/swrc/internal/account/domain/entity"

type AccountRepository interface {
	FindByID(id string) (*entity.Account, error)
	FindByEmail(email string) (*entity.Account, error)
	FindByGithubDisplayName(githubDisplayName string) (*entity.Account, error)
	Store(acc *entity.Account) (string, error)
	SaveVerification(acc *entity.Account) error
	SaveRole(acc *entity.Account) error
}
