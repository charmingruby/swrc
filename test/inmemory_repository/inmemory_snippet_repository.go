package inmemory_repository

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func NewInMemorySnippetRepository() *InMemorySnippetRepository {
	return &InMemorySnippetRepository{
		Items: []entity.Snippet{},
	}
}

type InMemorySnippetRepository struct {
	Items []entity.Snippet
}

func (r *InMemorySnippetRepository) Store(snippet entity.Snippet) error {
	r.Items = append(r.Items, snippet)
	return nil
}

func (r *InMemorySnippetRepository) FindByID(id string) (entity.Snippet, error) {
	for _, snp := range r.Items {
		if snp.ID == id {
			return snp, nil
		}
	}

	return entity.Snippet{}, core.NewNotFoundErr("snippet")
}
