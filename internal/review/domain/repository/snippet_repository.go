package repository

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type SnippetRepository interface {
	Store(snippet entity.Snippet) error
	FindByID(id string) (entity.Snippet, error)
}
