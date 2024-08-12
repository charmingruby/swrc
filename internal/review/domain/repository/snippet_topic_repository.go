package repository

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type SnippetTopicRepository interface {
	Store(topic entity.SnippetTopic) error
}
