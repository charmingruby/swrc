package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func (s *ReviewUseCaseRegistry) FetchCommentsUseCase(input dto.FetchCommentsInputDTO) (dto.FetchCommentsOutputDTO, error) {
	comments, err := s.CommentRepository.FindMany(
		input.ID,
		input.AccountID,
		input.SnippetTopicID,
		input.ParentCommentID,
	)
	if err != nil {
		return dto.FetchCommentsOutputDTO{
			Comments: []entity.Comment{},
		}, nil
	}

	return dto.FetchCommentsOutputDTO{
		Comments: comments,
	}, nil
}
