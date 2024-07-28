package review_dto

import "github.com/charmingruby/swrc/internal/review/domain/review_entity"

type GetSnippetTopicInputDTO struct {
	SnippetTopicID string
}

type GetSnippetTopicOutputDTO struct {
	SnippetTopic *review_entity.SnippetTopic
}
