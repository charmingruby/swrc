package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
)

const (
	modifySnippetTopicUseCase = "Modify Snippet Topic Use Case"
)

func (s *ReviewUseCaseRegistry) ModifySnippetTopicUseCase(input dto.ModifySnippetTopicInputDTO) error {
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

	nothingToChange := input.Title == topic.Title &&
		input.Description == topic.Description &&
		input.Status == topic.Status
	if nothingToChange {
		return core.NewNothingToChangeErr()
	}

	if input.Title != topic.Title {
		topic.Title = input.Title
	}

	if input.Description != topic.Description {
		topic.Description = input.Description
	}

	if input.Status != topic.Status {
		if err := topic.SetStatus(input.Status); err != nil {
			return err
		}
	}

	if err := s.SnippetTopicRepository.Save(topic); err != nil {
		logger.LogInternalErr(modifySnippetTopicUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
