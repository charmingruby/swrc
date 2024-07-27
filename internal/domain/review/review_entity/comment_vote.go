package review_entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/core"
)

func NewCommentVote(isUp bool, accountID, commentID string) (*CommentVote, error) {
	cv := CommentVote{
		ID:        core.NewID(),
		IsUp:      isUp,
		AccountID: accountID,
		CommentID: commentID,
		CreatedAt: time.Now(),
	}

	if err := core.ValidateStruct(cv); err != nil {
		return nil, err
	}

	return &cv, nil
}

type CommentVote struct {
	ID        string    `json:"id" validate:"required"`
	IsUp      bool      `json:"is_up" validate:"required"`
	AccountID string    `json:"account_id" validate:"required"`
	CommentID string    `json:"comment_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}
