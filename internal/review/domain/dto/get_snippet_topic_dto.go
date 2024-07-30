package dto

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type GetSnippetTopicInputDTO struct {
	SnippetTopicID string
}

type GetSnippetTopicOutputDTO struct {
	SnippetTopic *entity.SnippetTopic
}
