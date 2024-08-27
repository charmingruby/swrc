package factory

import (
	"github.com/charmingruby/swrc/internal/common/util"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/internal/review/domain/repository"
)

type MakeCommentVoteInput struct {
	IsUp      bool
	AccountID string
	CommentID string
}

func MakeCommentVote(
	repo repository.CommentVoteRepository,
	in MakeCommentVoteInput,
) (entity.CommentVote, error) {
	accountID := util.Ternary[string](in.AccountID == "", "invalid id", in.AccountID)
	commentID := util.Ternary[string](in.CommentID == "", "invalid id", in.CommentID)
	isUp := in.IsUp

	commentVote, err := entity.NewCommentVote(isUp, accountID, commentID)
	if err != nil {
		return entity.CommentVote{}, err
	}

	if err := repo.Store(*commentVote); err != nil {
		return entity.CommentVote{}, err
	}

	return *commentVote, nil
}
