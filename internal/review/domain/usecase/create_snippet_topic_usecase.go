package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

const (
	createSnippetTopicUseCase = "Create Snippet Topic Use Case"
)

func (s *ReviewUseCaseRegistry) CreateSnippetTopicUseCase(input dto.CreateSnippetTopicInputDTO) error {
	if err := s.AccountClient.ValidAccountExists(input.AccountID); err != nil {
		return err
	}

	snpTopic, err := entity.NewSnippetTopic(
		input.Title,
		input.Description,
		input.AccountID,
	)
	if err != nil {
		return err
	}

	if err := s.SnippetTopicRepository.Store(*snpTopic); err != nil {
		logger.LogInternalErr(createSnippetTopicUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
