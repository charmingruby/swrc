package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
)

const (
	chooseSnippetTopicBestAnswerUseCase = "Choose Snippet Topic Best Answer Use Case"
)

func (s *ReviewUseCaseRegistry) ChooseSnippetTopicBestAnswerUseCase(input dto.ChooseSnippetTopicBestAnswerInputDTO) error {
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

	comment, err := s.CommentRepository.FindByID(input.CommentID)
	if err != nil {
		return core.NewNotFoundErr("comment")
	}

	if comment.SnippetTopicID != topic.ID {
		return core.NewNotFoundErr("comment")
	}

	topic.BestAnswerID = comment.ID

	if err := s.SnippetTopicRepository.Save(topic); err != nil {
		logger.LogInternalErr(chooseSnippetTopicBestAnswerUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
