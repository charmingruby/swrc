package inmemory_repository

import (
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func NewInMemoryCommentRepository() *InMemoryCommentRepository {
	return &InMemoryCommentRepository{
		Items: []entity.Comment{},
	}
}

type InMemoryCommentRepository struct {
	Items []entity.Comment
}

func (r *InMemoryCommentRepository) FindByID(id string) (entity.Comment, error) {
	for _, comment := range r.Items {
		if comment.ID == id {
			return comment, nil
		}
	}

	return entity.Comment{}, ErrNotFound
}

func (r *InMemoryCommentRepository) Store(comment entity.Comment) error {
	r.Items = append(r.Items, comment)
	return nil
}
