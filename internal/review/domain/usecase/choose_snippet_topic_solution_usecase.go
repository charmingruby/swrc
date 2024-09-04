package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

const (
	chooseSnippetTopicSolutionUseCase = "Choose Snippet Topic Solution Use Case"
)

func (s *ReviewUseCaseRegistry) ChooseSnippetTopicSolutionUseCase(input dto.ChooseSnippetTopicSolutionInputDTO) error {
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

	snippet, err := s.SnippetRepository.FindByID(input.SnippetID)
	if err != nil {
		return core.NewNotFoundErr("snippet")
	}

	topic.SnippetSolutionID = snippet.ID
	topic.SetStatus(entity.SNIPPET_TOPIC_STATUS_CLOSED)

	if err := s.SnippetTopicRepository.Save(topic); err != nil {
		logger.LogInternalErr(chooseSnippetTopicSolutionUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
