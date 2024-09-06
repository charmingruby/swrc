package repository

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type CommentVoteRepository interface {
	FindByID(id string) (entity.CommentVote, error)
	FindByCommentIDAndAccountID(commentID, accountID string) (entity.CommentVote, error)
	Store(vote entity.CommentVote) error
	Delete(vote entity.CommentVote) error
}
