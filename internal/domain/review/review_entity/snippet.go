package review_entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/core"
)

const (
	SNIPPET_STATUS_OPEN   = "OPEN"
	SNIPPET_STATUS_CLOSED = "CLOSED"
)

func NewSnippet(title, description, codeSnippet, accountID string) (*Snippet, error) {
	snp := Snippet{
		ID:          core.NewID(),
		Title:       title,
		Description: description,
		CodeSnippet: codeSnippet,
		Status:      SNIPPET_STATUS_OPEN,
		Reactions:   0,
		Comments:    0,
		Votes:       0,
		AccountID:   accountID,
		CreatedAt:   time.Now(),
	}

	if err := core.ValidateStruct(snp); err != nil {
		return nil, err
	}

	return &snp, nil
}

type Snippet struct {
	ID          string    `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CodeSnippet string    `json:"code_snippet" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	Reactions   int       `json:"reactions"`
	Comments    int       `json:"comments"`
	Votes       int       `json:"votes"`
	AccountID   string    `json:"account_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at" validate:"required"`
}

func (snp *Snippet) Close() {
	snp.Status = SNIPPET_STATUS_CLOSED
}
