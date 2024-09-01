package dto

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type FetchSnippetTopicsInputDTO struct{}

type FetchSnippetTopicsOutputDTO struct {
	Topics []entity.SnippetTopic
}
