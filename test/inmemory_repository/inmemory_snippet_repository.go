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

func (r *InMemorySnippetRepository) DeleteManyByTopicID(topicID string) error {
	remainingVotes := []entity.Snippet{}

	for _, c := range r.Items {
		if c.TopicID != topicID {
			remainingVotes = append(remainingVotes, c)
		}
	}

	r.Items = remainingVotes

	return nil
}
