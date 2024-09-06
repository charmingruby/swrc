package inmemory_repository

import (
	"github.com/charmingruby/swrc/internal/common/core"
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

func (r *InMemoryCommentRepository) FindMany(id string, accountID string, snippetTopicID string, parentCommentID string) ([]entity.Comment, error) {
	var results []entity.Comment

	for _, item := range r.Items {
		idMatch := id == "" || item.ID == id

		statusMatch := accountID == "" || item.AccountID == accountID

		snippetTopicIDMatch := snippetTopicID == "" || item.SnippetTopicID == snippetTopicID

		parentCommentIDMatch := parentCommentID == "" || item.ParentCommentID == parentCommentID

		if idMatch && statusMatch && snippetTopicIDMatch && parentCommentIDMatch {
			results = append(results, item)
		}
	}

	if len(results) == 0 {
		return nil, core.NewNotFoundErr("comments")
	}

	return results, nil
}

func (r *InMemoryCommentRepository) Store(comment entity.Comment) error {
	r.Items = append(r.Items, comment)
	return nil
}

func (r *InMemoryCommentRepository) Delete(comment entity.Comment) error {
	for idx, tpc := range r.Items {
		if tpc.ID == comment.ID {
			r.Items = append(r.Items[:idx], r.Items[idx+1:]...)
			return nil
		}
	}

	return core.NewNotFoundErr("comment")
}

func (r *InMemoryCommentRepository) DeleteManyByParentCommentID(parentCommentID string) error {
	var collectCommentsToDelete func(parentID string) []string

	collectCommentsToDelete = func(parentID string) []string {
		idsToDelete := []string{parentID}

		for _, c := range r.Items {
			if c.ParentCommentID == parentID {
				idsToDelete = append(idsToDelete, collectCommentsToDelete(c.ID)...)
			}
		}

		return idsToDelete
	}

	idsToDelete := collectCommentsToDelete(parentCommentID)

	remainingComments := []entity.Comment{}
	for _, c := range r.Items {
		shouldDelete := false
		for _, id := range idsToDelete {
			if c.ID == id {
				shouldDelete = true
				break
			}
		}

		if !shouldDelete {
			remainingComments = append(remainingComments, c)
		}
	}

	r.Items = remainingComments

	return nil
}
