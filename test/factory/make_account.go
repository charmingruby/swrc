package factory

import (
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/account/domain/repository"
	"github.com/charmingruby/swrc/internal/common/util"
	"github.com/charmingruby/swrc/test/fake"
)

type MakeAccountInput struct {
	GithubDisplayName string
	Email             string
	Password          string
	IsValid           bool
	Verified          bool
	Role              string
}

func MakeAccount(
	repo repository.AccountRepository,
	in MakeAccountInput) (*entity.Account, error) {
	ghDisplayName := util.Ternary[string](in.GithubDisplayName == "", "charmingruby", in.GithubDisplayName)
	email := util.Ternary[string](in.Email == "", "dummy@example.com", in.Email)
	password := util.Ternary[string](in.Password == "", "password123", in.Password)
	isValid := in.IsValid
	verified := in.Verified
	role := util.Ternary[string](in.Role == "", entity.ACCOUNT_ROLE_DEVELOPER, in.Role)

	hashedPassword, _ := fake.NewFakeHashService().GenerateHash(password)
	acc, err := entity.NewAccount(
		ghDisplayName,
		email,
		password,
	)
	if err != nil {
		return nil, err
	}
	acc.Password = hashedPassword

	acc.Verification.Verified = verified
	acc.Verification.IsValid = isValid
	acc.Role = role

	if err := repo.Store(*acc); err != nil {
		return nil, err
	}

	return acc, nil
}
