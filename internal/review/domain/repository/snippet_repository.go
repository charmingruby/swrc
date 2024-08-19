package repository

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type SnippetRepository interface {
	Store(snippet entity.Snippet) error
	Save(snippet entity.Snippet) error
	FindByID(id string) (entity.Snippet, error)
	FindManyByTopicID(topicID string) ([]entity.Snippet, error)
	DeleteMany([]entity.Snippet) error
}
