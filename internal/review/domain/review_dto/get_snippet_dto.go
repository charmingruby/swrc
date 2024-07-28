package review_dto

import "github.com/charmingruby/swrc/internal/review/domain/review_entity"

type GetSnippetInputDTO struct {
	SnippetID string
}

type GetSnippetOutputDTO struct {
	Snippet *review_entity.Snippet
}
