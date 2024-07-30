package entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/common/core"
)

func NewComment(content, accountID, snippetID, parentCommentID string) (*Comment, error) {
	c := Comment{
		ID:              core.NewID(),
		Content:         content,
		Votes:           0,
		AccountID:       accountID,
		ParentCommentID: parentCommentID,
		SnippetID:       snippetID,
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
	Votes           int       `json:"votes"`
	AccountID       string    `json:"account_id" validate:"required"`
	ParentCommentID string    `json:"parent_comment_id"`
	SnippetID       string    `json:"snippet_id" validate:"required"`
	CreatedAt       time.Time `json:"created_at" validate:"required"`
}
