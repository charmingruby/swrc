package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func (s *ReviewUseCaseRegistry) FetchSnippetTopicsUseCase(input dto.FetchSnippetTopicsInputDTO) (dto.FetchSnippetTopicsOutputDTO, error) {
	topics, err := s.SnippetTopicRepository.FindMany(input.ID, input.Status, input.AccountID)
	if err != nil {
		return dto.FetchSnippetTopicsOutputDTO{
			Topics: []entity.SnippetTopic{},
		}, nil
	}

	return dto.FetchSnippetTopicsOutputDTO{
		Topics: topics,
	}, nil
}
