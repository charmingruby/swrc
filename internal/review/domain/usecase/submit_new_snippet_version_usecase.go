package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

const (
	submitNewSnippetVersionUseCase = "Submit new Snippet version Use Case"
)

func (s *ReviewUseCaseRegistry) SubmitNewSnippetVersionUseCase(input dto.SubmitNewSnippetVersionInputDTO) error {
	if err := s.AccountClient.ValidAccountExists(input.AccountID); err != nil {
		return err
	}

	topic, err := s.SnippetTopicRepository.FindByID(input.TopicID)
	if err != nil {
		return core.NewNotFoundErr("snippet topic")
	}

	if topic.AccountID != input.AccountID {
		return core.NewUnauthorizedErr()
	}

	newVersion := topic.CurrentVersion + 1

	snippet, err := entity.NewSnippet(newVersion, input.CodeSnippet, input.Message, input.TopicID)
	if err != nil {
		return err
	}

	if err := s.SnippetRepository.Store(*snippet); err != nil {
		logger.LogInternalErr(submitNewSnippetVersionUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
