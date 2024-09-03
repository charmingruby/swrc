package repository

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type SnippetTopicRepository interface {
	FindByID(id string) (entity.SnippetTopic, error)
	FindMany(id, status, accountID string) ([]entity.SnippetTopic, error)
	Delete(id string) error
	Store(topic entity.SnippetTopic) error
	Save(topic entity.SnippetTopic) error
}
