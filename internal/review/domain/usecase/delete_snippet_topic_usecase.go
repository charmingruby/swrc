package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
)

const (
	deleteSnippetTopicUseCase = "Delete Snippet Topic Use Case"
)

func (s *ReviewUseCaseRegistry) DeleteSnippetTopicUseCase(input dto.DeleteSnippetTopicInputDTO) error {
	if err := s.AccountClient.ValidAccountExists(input.AccountID); err != nil {
		return err
	}

	topic, err := s.SnippetTopicRepository.FindByID(input.SnippetTopicID)
	if err != nil {
		return core.NewNotFoundErr("snippet topic")
	}

	if topic.AccountID != input.AccountID {
		return core.NewUnauthorizedErr()
	}

	snippetsToDelete, err := s.SnippetRepository.FindManyByTopicID(topic.ID)
	if err != nil {
		logger.LogInternalErr(deleteSnippetTopicUseCase, err)
		return core.NewInternalErr()
	}

	if len(snippetsToDelete) > 0 {
		if err := s.SnippetRepository.DeleteMany(snippetsToDelete); err != nil {
			logger.LogInternalErr(deleteSnippetTopicUseCase, err)
			return core.NewInternalErr()
		}
	}

	if err := s.SnippetTopicRepository.Delete(topic.ID); err != nil {
		logger.LogInternalErr(deleteSnippetTopicUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
