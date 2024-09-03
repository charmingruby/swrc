package dto

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type FetchSnippetInputDTO struct {
	SnippetTopicID string
}

type FetchSnippetOutputDTO struct {
	Snippets []entity.Snippet
}
