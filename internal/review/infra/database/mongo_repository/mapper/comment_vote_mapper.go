package mapper

import (
	"time"

	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func DomainCommentVoteToMongo(vote entity.CommentVote) MongoCommentVote {
	return MongoCommentVote{
		ID:        vote.ID,
		IsUp:      vote.IsUp,
		CommentID: vote.CommentID,
		AccountID: vote.AccountID,
		CreatedAt: vote.CreatedAt,
	}
}

func MongoCommentVoteToDomain(vote MongoCommentVote) entity.CommentVote {
	return entity.CommentVote{
		ID:        vote.ID,
		IsUp:      vote.IsUp,
		CommentID: vote.CommentID,
		AccountID: vote.AccountID,
		CreatedAt: vote.CreatedAt,
	}
}

type MongoCommentVote struct {
	ID        string    `json:"id" bson:"_id"`
	IsUp      bool      `json:"is_up" bson:"is_up"`
	AccountID string    `json:"account_id" bson:"account_id"`
	CommentID string    `json:"comment_id" bson:"comment_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
