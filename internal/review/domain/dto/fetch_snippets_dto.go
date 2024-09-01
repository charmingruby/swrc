package dto

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type FetchSnippetInputDTO struct{}

type FetchSnippetOutputDTO struct {
	Snippets []entity.Snippet
}
