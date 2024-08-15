package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
)

func (s *ReviewUseCaseRegistry) GetSnippetUseCase(input dto.GetSnippetInputDTO) (dto.GetSnippetOutputDTO, error) {
	snippet, err := s.SnippetRepository.FindByID(input.SnippetID)
	if err != nil {
		return dto.GetSnippetOutputDTO{}, core.NewNotFoundErr("snippet")
	}

	return dto.GetSnippetOutputDTO{
		Snippet: &snippet,
	}, nil
}
