package mapper

import (
	"time"

	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/account/domain/entity/account_valueobj"
)

func DomainAccountToMongo(acc entity.Account) MongoAccount {
	return MongoAccount{
		ID:                acc.ID,
		GithubDisplayName: acc.GithubDisplayName,
		Email:             acc.Email,
		Password:          acc.Password,
		IsValid:           acc.Verification.IsValid,
		Verified:          acc.Verification.Verified,
		Role:              acc.Role,
		CreatedAt:         acc.CreatedAt,
	}
}

func MongoAccountToDomain(acc MongoAccount) entity.Account {
	return entity.Account{
		ID:                acc.ID,
		GithubDisplayName: acc.GithubDisplayName,
		Email:             acc.Email,
		Password:          acc.Password,
		Verification: &account_valueobj.Verification{
			IsValid:  acc.IsValid,
			Verified: acc.Verified,
		},
		Role:      acc.Role,
		CreatedAt: acc.CreatedAt,
	}
}

type MongoAccount struct {
	ID                string    `json:"id" bson:"_id"`
	GithubDisplayName string    `json:"github_display_name" bson:"github_display_name"`
	Email             string    `json:"email" bson:"email"`
	Password          string    `json:"password" bson:"password"`
	IsValid           bool      `json:"is_valid" bson:"is_valid"`
	Verified          bool      `json:"verified" bson:"verified"`
	Role              string    `json:"role" bson:"role"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
}
