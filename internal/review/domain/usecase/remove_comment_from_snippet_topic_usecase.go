package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
)

const (
	removeCommentFromTopicUseCase = "Remove Comment From Topic Use Case"
)

func (r *ReviewUseCaseRegistry) RemoveCommentFromSnippetTopicUseCase(input dto.RemoveCommentFromSnippetTopicInputDTO) error {
	if err := r.AccountClient.ValidAccountExists(input.AccountID); err != nil {
		return err
	}

	comment, err := r.CommentRepository.FindByID(input.CommentID)
	if err != nil {
		return core.NewNotFoundErr("comment")
	}

	if comment.AccountID != input.AccountID {
		return core.NewUnauthorizedErr()
	}

	if err := r.CommentVoteRepository.DeleteManyByCommentID(comment.ID); err != nil {
		logger.LogInternalErr(removeCommentFromTopicUseCase, err)
		return core.NewInternalErr()
	}

	if err := r.CommentRepository.Delete(comment); err != nil {
		logger.LogInternalErr(removeCommentFromTopicUseCase, err)
		return core.NewInternalErr()
	}

	if err := r.CommentRepository.DeleteManyByParentCommentID(input.CommentID); err != nil {
		logger.LogInternalErr(removeCommentFromTopicUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
