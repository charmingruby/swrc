package inmemory_repository

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func NewInMemoryCommentVoteRepository() *InMemoryCommentVoteRepository {
	return &InMemoryCommentVoteRepository{
		Items: []entity.CommentVote{},
	}
}

type InMemoryCommentVoteRepository struct {
	Items []entity.CommentVote
}

func (r *InMemoryCommentVoteRepository) Store(cv entity.CommentVote) error {
	r.Items = append(r.Items, cv)
	return nil
}

func (r *InMemoryCommentVoteRepository) Delete(e entity.CommentVote) error {
	for idx, cv := range r.Items {
		if cv.ID == e.ID {
			r.Items = append(r.Items[:idx], r.Items[idx+1:]...)
			return nil
		}
	}

	return core.NewNotFoundErr("comment vote")
}

func (r *InMemoryCommentVoteRepository) FindByID(id string) (entity.CommentVote, error) {
	for _, cv := range r.Items {
		if cv.ID == id {
			return cv, nil
		}
	}

	return entity.CommentVote{}, core.NewNotFoundErr("comment vote")
}

func (r *InMemoryCommentVoteRepository) DeleteManyByCommentID(commentID string) error {
	remainingVotes := []entity.CommentVote{}

	for _, cv := range r.Items {
		if cv.CommentID != commentID {
			remainingVotes = append(remainingVotes, cv)
		}
	}

	r.Items = remainingVotes

	return nil
}

func (r *InMemoryCommentVoteRepository) FindByCommentIDAndAccountID(commentID, accountID string) (entity.CommentVote, error) {
	for _, cv := range r.Items {
		if cv.CommentID == commentID && cv.AccountID == accountID {
			return cv, nil
		}
	}

	return entity.CommentVote{}, core.NewNotFoundErr("comment vote")
}
