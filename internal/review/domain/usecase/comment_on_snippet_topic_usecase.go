package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

const (
	commentOnSnippetTopicUseCase = "Comment On Snippet Topic Use Case"
)

func (s *ReviewUseCaseRegistry) CommentOnSnippetTopicUseCase(input dto.CommentOnSnippetTopicInputDTO) error {
	if err := s.AccountClient.ValidAccountExists(input.AccountID); err != nil {
		return err
	}

	_, err := s.SnippetTopicRepository.FindByID(input.SnippetTopicID)
	if err != nil {
		return core.NewNotFoundErr("snippet topic")
	}

	if input.ParentCommentID != "" {
		if _, err := s.CommentRepository.FindByID(input.ParentCommentID); err != nil {
			return core.NewNotFoundErr("parent comment")
		}
	}

	comment, err := entity.NewComment(input.Content, input.AccountID, input.SnippetTopicID, input.ParentCommentID)
	if err != nil {
		return err
	}

	if err := s.CommentRepository.Store(*comment); err != nil {
		logger.LogInternalErr(commentOnSnippetTopicUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
