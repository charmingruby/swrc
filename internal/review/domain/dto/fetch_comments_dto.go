package dto

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type FetchCommentsInputDTO struct {
	ID              string
	AccountID       string
	SnippetTopicID  string
	ParentCommentID string
}

type FetchCommentsOutputDTO struct {
	Comments []entity.Comment
}
