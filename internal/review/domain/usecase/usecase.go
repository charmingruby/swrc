package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/adapter"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/repository"
)

type ReviewUseCase interface {
	CreateSnippetTopicUseCase(dto dto.CreateSnippetTopicInputDTO) error
	SubmitNewSnippetVersionUseCase(dto dto.SubmitNewSnippetVersionInputDTO) error
	ChooseSnippetTopicSolutionUseCase(dto dto.ChooseSnippetTopicSolutionInputDTO) error
	GetSnippetTopicUseCase(dto dto.GetSnippetTopicInputDTO) (dto.GetSnippetTopicOutputDTO, error)
	GetSnippetUseCase(dto dto.GetSnippetInputDTO) (dto.GetSnippetOutputDTO, error)
	ModifySnippetTopicUseCase(dto dto.ModifySnippetTopicInputDTO) error
	ModifySnippetUseCase(dto dto.ModifySnippetInputDTO) error
	DeleteSnippetTopicUseCase(dto dto.DeleteSnippetTopicInputDTO) error
	ChooseSnippetTopicBestAnswerUseCase(dto dto.ChooseSnippetTopicBestAnswerInputDTO) error
	CommentOnSnippetTopicUseCase(dto dto.CommentOnSnippetTopicInputDTO) error
	RemoveCommentFromSnippetTopicUseCase(dto dto.RemoveCommentFromSnippetTopicInputDTO) error
	VoteOnCommentUseCase(dto dto.VoteOnCommentInputDTO) error
	RemoveVoteFromCommentUseCase(dto dto.RemoveVoteFromCommentInputDTO) error
	FetchSnippetTopicsUseCase(dto dto.FetchSnippetTopicsInputDTO) (dto.FetchSnippetTopicsOutputDTO, error)
	FetchSnippetsUseCase(dto dto.FetchSnippetInputDTO) (dto.FetchSnippetOutputDTO, error)
	FetchCommentsUseCase(dto dto.FetchCommentsInputDTO) (dto.FetchCommentsOutputDTO, error)
}

func NewReviewUseCaseRegistry(
	snippetRepository repository.SnippetRepository,
	snippetTopicRepository repository.SnippetTopicRepository,
	commentRepository repository.CommentRepository,
	commentVoteRepository repository.CommentVoteRepository,
	accountClient adapter.AccountClient,
) *ReviewUseCaseRegistry {
	return &ReviewUseCaseRegistry{
		SnippetRepository:      snippetRepository,
		SnippetTopicRepository: snippetTopicRepository,
		CommentRepository:      commentRepository,
		CommentVoteRepository:  commentVoteRepository,
		AccountClient:          accountClient,
	}
}

type ReviewUseCaseRegistry struct {
	SnippetRepository      repository.SnippetRepository
	SnippetTopicRepository repository.SnippetTopicRepository
	CommentRepository      repository.CommentRepository
	CommentVoteRepository  repository.CommentVoteRepository
	AccountClient          adapter.AccountClient
}
