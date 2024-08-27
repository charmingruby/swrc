package repository

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type CommentRepository interface {
	FindByID(id string) (entity.Comment, error)
	Store(comment entity.Comment) error
	Save(comment entity.Comment) error
	Delete(comment entity.Comment) error
	DeleteManyByParentCommentID(parentCommentID string) error
}
