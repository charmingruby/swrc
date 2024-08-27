package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

const (
	voteOnCommentUseCase = "Vote On Comment Use Case"
)

func (r *ReviewUseCaseRegistry) VoteOnCommentUseCase(input dto.VoteOnCommentInputDTO) error {
	if err := r.AccountClient.ValidAccountExists(input.AccountID); err != nil {
		return err
	}

	comment, err := r.CommentRepository.FindByID(input.CommentID)
	if err != nil {
		return core.NewNotFoundErr("comment")
	}

	if _, err := r.CommentVoteRepository.FindByCommentIDAndAccountID(input.CommentID, input.AccountID); err == nil {
		return core.NewAlreadyExistsErr("comment vote")
	}

	if _, err := entity.NewCommentVote(input.IsUp, input.AccountID, input.CommentID); err != nil {
		return err
	}

	topic, err := r.SnippetTopicRepository.FindByID(comment.SnippetTopicID)
	if err != nil {
		return core.NewNotFoundErr("snippet topic")
	}

	topic.Votes += 1
	if err := r.SnippetTopicRepository.Save(topic); err != nil {
		logger.LogInternalErr(voteOnCommentUseCase, err)
		return core.NewInternalErr()
	}

	comment.Votes += 1
	if err := r.CommentRepository.Save(comment); err != nil {
		logger.LogInternalErr(voteOnCommentUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
