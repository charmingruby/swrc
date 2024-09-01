package entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/common/core"
)

func NewComment(content, accountID, snippetTopicID, parentCommentID string) (*Comment, error) {
	c := Comment{
		ID:              core.NewID(),
		Content:         content,
		AccountID:       accountID,
		ParentCommentID: parentCommentID,
		SnippetTopicID:  snippetTopicID,
		CreatedAt:       time.Now(),
	}

	if err := core.ValidateStruct(c); err != nil {
		return nil, err
	}

	return &c, nil
}

type Comment struct {
	ID              string    `json:"id" validate:"required"`
	Content         string    `json:"content" validate:"required"`
	AccountID       string    `json:"account_id" validate:"required"`
	ParentCommentID string    `json:"parent_comment_id"`
	SnippetTopicID  string    `json:"snippet_topic_id" validate:"required"`
	CreatedAt       time.Time `json:"created_at" validate:"required"`
}
