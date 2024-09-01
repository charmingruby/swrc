package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func (r *ReviewUseCaseRegistry) VoteOnCommentUseCase(input dto.VoteOnCommentInputDTO) error {
	if err := r.AccountClient.ValidAccountExists(input.AccountID); err != nil {
		return err
	}

	_, err := r.CommentRepository.FindByID(input.CommentID)
	if err != nil {
		return core.NewNotFoundErr("comment")
	}

	if _, err := r.CommentVoteRepository.FindByCommentIDAndAccountID(input.CommentID, input.AccountID); err == nil {
		return core.NewAlreadyExistsErr("comment vote")
	}

	if _, err := entity.NewCommentVote(input.IsUp, input.AccountID, input.CommentID); err != nil {
		return err
	}

	return nil
}
