package review_entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/common/core"
)

func NewDeveloperMetrics(accountID string) (*DeveloperMetrics, error) {
	dm := DeveloperMetrics{
		ID:                core.NewID(),
		SnippetsPublished: 0,
		Stars:             0,
		AccountID:         accountID,
		CreatedAt:         time.Now(),
	}

	if err := core.ValidateStruct(dm); err != nil {
		return nil, err
	}

	return &dm, nil
}

type DeveloperMetrics struct {
	ID                string    `json:"id" validate:"required"`
	SnippetsPublished int       `json:"snippets_published"`
	Stars             int       `json:"stars"` // snippets resolved by this account
	AccountID         string    `json:"account_id" validate:"required"`
	CreatedAt         time.Time `json:"created_at" validate:"required"`
}
