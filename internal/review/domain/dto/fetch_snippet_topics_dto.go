package dto

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type FetchSnippetTopicsInputDTO struct {
	ID        string
	Status    string
	AccountID string
}

type FetchSnippetTopicsOutputDTO struct {
	Topics []entity.SnippetTopic
}
