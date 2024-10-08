package inmemory_repository

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

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

func (r *InMemorySnippetTopicRepository) Save(topic entity.SnippetTopic) error {
	for idx, tpc := range r.Items {
		if tpc.ID == topic.ID {
			r.Items[idx] = topic
			return nil
		}
	}

	return core.NewNotFoundErr("snippet topic")

}

func (r *InMemorySnippetTopicRepository) Delete(id string) error {
	for idx, tpc := range r.Items {
		if tpc.ID == id {
			r.Items = append(r.Items[:idx], r.Items[idx+1:]...)
			return nil
		}
	}

	return core.NewNotFoundErr("snippet topic")
}

func (r *InMemorySnippetTopicRepository) FindByID(ID string) (entity.SnippetTopic, error) {
	for _, tpc := range r.Items {
		if tpc.ID == ID {
			return tpc, nil
		}
	}

	return entity.SnippetTopic{}, core.NewNotFoundErr("snippet topic")
}

func (r *InMemorySnippetTopicRepository) FindMany(id, status, accountID string) ([]entity.SnippetTopic, error) {
	var results []entity.SnippetTopic

	for _, item := range r.Items {
		idMatch := id == "" || item.ID == id

		statusMatch := status == "" || item.Status == status

		accountIDMatch := accountID == "" || item.AccountID == accountID

		if idMatch && statusMatch && accountIDMatch {
			results = append(results, item)
		}
	}

	if len(results) == 0 {
		return nil, core.NewNotFoundErr("snippet topics")
	}

	return results, nil
}
