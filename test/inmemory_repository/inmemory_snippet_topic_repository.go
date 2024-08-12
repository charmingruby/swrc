package inmemory_repository

import "github.com/charmingruby/swrc/internal/review/domain/entity"

func NewInMemorySnippetTopicRepository() *InMemorySnippetTopicRepository {
	return &InMemorySnippetTopicRepository{
		Items: []entity.SnippetTopic{},
	}
}

type InMemorySnippetTopicRepository struct {
	Items []entity.SnippetTopic
}

func (r *InMemorySnippetTopicRepository) Store(topic entity.SnippetTopic) error {
	r.Items = append(r.Items, topic)
	return nil
}
