package review_entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/common/core"
)

func NewSnippet(version int, codeSnippet, message, topicID string) (*Snippet, error) {
	snp := Snippet{
		ID:          core.NewID(),
		Version:     version,
		CodeSnippet: codeSnippet,
		Message:     message,
		TopicID:     topicID,
		CreatedAt:   time.Now(),
	}

	if err := core.ValidateStruct(snp); err != nil {
		return nil, err
	}

	return &snp, nil
}

type Snippet struct {
	ID          string    `json:"id" validate:"required"`
	Version     int       `json:"version" validate:"required"`
	CodeSnippet string    `json:"code_snippet" validate:"required"`
	Message     string    `json:"message"`
	TopicID     string    `json:"topic_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at" validate:"required"`
}
