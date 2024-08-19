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

func (r *InMemorySnippetRepository) Save(snippet entity.Snippet) error {
	for idx, snp := range r.Items {
		if snp.ID == snippet.ID {
			r.Items[idx] = snippet
			return nil
		}
	}

	return core.NewNotFoundErr("snippet")
}

func (r *InMemorySnippetRepository) FindByID(id string) (entity.Snippet, error) {
	for _, snp := range r.Items {
		if snp.ID == id {
			return snp, nil
		}
	}

	return entity.Snippet{}, core.NewNotFoundErr("snippet")
}

func (r *InMemorySnippetRepository) FindManyByTopicID(topicID string) ([]entity.Snippet, error) {
	snippets := []entity.Snippet{}

	for _, snp := range r.Items {
		if snp.TopicID == topicID {
			snippets = append(snippets, snp)
		}
	}

	return snippets, nil
}

func (r *InMemorySnippetRepository) DeleteMany(snippets []entity.Snippet) error {
	remainingSnippets := []entity.Snippet{}

	toBeRemoved := make(map[string]struct{})
	for _, snippet := range snippets {
		toBeRemoved[snippet.ID] = struct{}{}
	}

	for _, s := range r.Items {
		if _, found := toBeRemoved[s.ID]; !found {
			remainingSnippets = append(remainingSnippets, s)
		}
	}

	r.Items = remainingSnippets

	return nil
}
