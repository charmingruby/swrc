package entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/common/core"
)

const (
	SNIPPET_TOPIC_STATUS_OPEN   = "OPEN"
	SNIPPET_TOPIC_STATUS_CLOSED = "CLOSED"
)

func NewSnippetTopic(title, description, accountID string) (*SnippetTopic, error) {
	snp := SnippetTopic{
		ID:             core.NewID(),
		Title:          title,
		Description:    description,
		Status:         SNIPPET_TOPIC_STATUS_OPEN,
		BestAnswerID:   "",
		CurrentVersion: 0,
		Comments:       0,
		Votes:          0,
		AccountID:      accountID,
		CreatedAt:      time.Now(),
	}

	if err := core.ValidateStruct(snp); err != nil {
		return nil, err
	}

	return &snp, nil
}

type SnippetTopic struct {
	ID                string    `json:"id" validate:"required"`
	Title             string    `json:"title" validate:"required"`
	Description       string    `json:"description" validate:"required"`
	Status            string    `json:"status" validate:"required"`
	CurrentVersion    int       `json:"current_version"`
	BestAnswerID      string    `json:"best_answer_id"`
	SnippetSolutionID string    `json:"snippet_solution_id"`
	Reactions         int       `json:"reactions"`
	Comments          int       `json:"comments"`
	Votes             int       `json:"votes"`
	AccountID         string    `json:"account_id" validate:"required"`
	CreatedAt         time.Time `json:"created_at" validate:"required"`
}

func (snp *SnippetTopic) Close() {
	snp.Status = SNIPPET_TOPIC_STATUS_CLOSED
}
