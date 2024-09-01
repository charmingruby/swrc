package dto

import "github.com/charmingruby/swrc/internal/review/domain/entity"

type FetchCommentsInputDTO struct{}

type FetchCommentsOutputDTO struct {
	Comments []entity.Comment
}
