package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
)

const (
	modifySnippetUseCase = "Modify Snippet Use Case"
)

func (s *ReviewUseCaseRegistry) ModifySnippetUseCase(input dto.ModifySnippetInputDTO) error {
	err := s.AccountClient.ValidAccountExists(input.AccountID)
	if err != nil {
		return err
	}

	snippet, err := s.SnippetRepository.FindByID(input.SnippetID)
	if err != nil {
		return core.NewNotFoundErr("snippet")
	}

	topic, err := s.SnippetTopicRepository.FindByID(snippet.TopicID)
	if err != nil {
		return core.NewNotFoundErr("snippet topic")
	}

	if topic.AccountID != input.AccountID {
		return core.NewUnauthorizedErr()
	}

	if input.CodeSnippet == snippet.CodeSnippet && input.Message == snippet.Message {
		return core.NewNothingToChangeErr()
	}

	if input.CodeSnippet != snippet.CodeSnippet {
		snippet.CodeSnippet = input.CodeSnippet
	}

	if input.Message != snippet.Message {
		snippet.Message = input.Message
	}

	if err := s.SnippetRepository.Save(snippet); err != nil {
		logger.LogInternalErr(modifySnippetUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
