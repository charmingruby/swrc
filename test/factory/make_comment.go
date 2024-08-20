package factory

import (
	"github.com/charmingruby/swrc/internal/common/util"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/internal/review/domain/repository"
)

type MakeCommentInput struct {
	Content         string
	Votes           int
	AccountID       string
	ParentCommentID string
	SnippetTopicID  string
}

func MakeComment(
	repo repository.CommentRepository,
	in MakeCommentInput,
) (entity.Comment, error) {
	content := util.Ternary[string](in.Content == "", "dummy content", in.Content)
	votes := util.Ternary[int](in.Votes == 0, 0, in.Votes)
	accountID := util.Ternary[string](in.AccountID == "", "invalid id", in.AccountID)
	parentCommentID := util.Ternary[string](in.ParentCommentID == "", "invalid id", in.ParentCommentID)
	snippetTopicID := util.Ternary[string](in.SnippetTopicID == "", "invalid id", in.SnippetTopicID)

	comment, err := entity.NewComment(content, accountID, snippetTopicID, parentCommentID)
	if err != nil {
		return entity.Comment{}, err
	}

	comment.Votes = votes

	if err := repo.Store(*comment); err != nil {
		return entity.Comment{}, err
	}

	return *comment, nil
}
