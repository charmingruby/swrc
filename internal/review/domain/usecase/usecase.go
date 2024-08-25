package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/adapter"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/repository"
)

type ReviewUseCase interface {
	// # -
	// SNIPPETS
	// # -
	CreateSnippetTopicUseCase(dto dto.CreateSnippetTopicInputDTO) error
	SubmitNewSnippetVersionUseCase(dto dto.SubmitNewSnippetVersionInputDTO) error
	ChooseSnippetTopicSolutionUseCase(dto dto.ChooseSnippetTopicSolutionInputDTO) error
	GetSnippetTopicUseCase(dto dto.GetSnippetTopicInputDTO) (dto.GetSnippetTopicOutputDTO, error)
	GetSnippetUseCase(dto dto.GetSnippetInputDTO) (dto.GetSnippetOutputDTO, error)
	ModifySnippetTopicUseCase(dto dto.ModifySnippetTopicInputDTO) error
	ModifySnippetUseCase(dto dto.ModifySnippetInputDTO) error
	DeleteSnippetTopicUseCase(dto dto.DeleteSnippetTopicInputDTO) error

	// # -
	// COMMENTS AND VOTES
	// # -
	ChooseSnippetTopicBestAnswerUseCase(dto dto.ChooseSnippetTopicBestAnswerInputDTO) error
	CommentOnSnippetTopicUseCase(dto dto.CommentOnSnippetTopicInputDTO) error
	RemoveCommentFromSnippetTopicUseCase(dto dto.RemoveCommentFromSnippetTopicInputDTO) error
	// VoteCommentUseCase()
	// RemoveVoteFromCommentUseCase()

	// # -
	// LIST
	// # -
	// FetchOpenSnippetsUseCase()
	// FetchAccountSnippetsUseCase()
	// FetchClosedSnippetsUseCase()
}

func NewReviewUseCaseRegistry(
	snippetRepository repository.SnippetRepository,
	snippetTopicRepository repository.SnippetTopicRepository,
	commentRepository repository.CommentRepository,
	accountClient adapter.AccountClient,
) *ReviewUseCaseRegistry {
	return &ReviewUseCaseRegistry{
		SnippetRepository:      snippetRepository,
		SnippetTopicRepository: snippetTopicRepository,
		CommentRepository:      commentRepository,
		AccountClient:          accountClient,
	}
}

type ReviewUseCaseRegistry struct {
	SnippetRepository      repository.SnippetRepository
	SnippetTopicRepository repository.SnippetTopicRepository
	CommentRepository      repository.CommentRepository
	AccountClient          adapter.AccountClient
}
