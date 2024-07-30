package dto

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type GetSnippetInputDTO struct {
	SnippetID string
}

type GetSnippetOutputDTO struct {
	Snippet *entity.Snippet
}
