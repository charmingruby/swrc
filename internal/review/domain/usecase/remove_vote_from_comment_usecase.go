package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
)

const (
	removeVoteFromCommentUseCase = "Remove Vote From Comment Use Case"
)

func (r *ReviewUseCaseRegistry) RemoveVoteFromCommentUseCase(input dto.RemoveVoteFromCommentInputDTO) error {
	if err := r.AccountClient.ValidAccountExists(input.AccountID); err != nil {
		return err
	}

	vote, err := r.CommentVoteRepository.FindByID(input.CommentVoteID)
	if err != nil {
		return core.NewNotFoundErr("comment vote")
	}

	isOwner := vote.AccountID == input.AccountID
	if !isOwner {
		return core.NewUnauthorizedErr()
	}

	if err := r.CommentVoteRepository.Delete(vote); err != nil {
		logger.LogInternalErr(removeVoteFromCommentUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
