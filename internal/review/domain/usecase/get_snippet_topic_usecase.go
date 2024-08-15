package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
)

func (s *ReviewUseCaseRegistry) GetSnippetTopicUseCase(input dto.GetSnippetTopicInputDTO) (dto.GetSnippetTopicOutputDTO, error) {
	topic, err := s.SnippetTopicRepository.FindByID(input.SnippetTopicID)
	if err != nil {
		return dto.GetSnippetTopicOutputDTO{}, core.NewNotFoundErr("snippet topic")
	}

	return dto.GetSnippetTopicOutputDTO{
		SnippetTopic: &topic,
	}, nil
}
