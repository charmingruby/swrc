package repository

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type SnippetTopicRepository interface {
	FindByID(id string) (entity.SnippetTopic, error)
	Store(topic entity.SnippetTopic) error
}
