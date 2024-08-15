package usecase

import "github.com/charmingruby/swrc/internal/review/domain/dto"

func (s *ReviewUseCaseRegistry) GetSnippetUseCase(input dto.GetSnippetInputDTO) (dto.GetSnippetOutputDTO, error) {
	return dto.GetSnippetOutputDTO{}, nil
}
