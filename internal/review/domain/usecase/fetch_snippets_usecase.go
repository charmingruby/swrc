package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func (s *ReviewUseCaseRegistry) FetchSnippetsUseCase(input dto.FetchSnippetInputDTO) (dto.FetchSnippetOutputDTO, error) {
	snippets, err := s.SnippetRepository.FindManyByTopicID(input.SnippetTopicID)
	if err != nil {
		return dto.FetchSnippetOutputDTO{
			Snippets: []entity.Snippet{},
		}, nil
	}

	return dto.FetchSnippetOutputDTO{
		Snippets: snippets,
	}, nil
}
