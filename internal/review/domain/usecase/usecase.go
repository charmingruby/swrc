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
	ChooseSnippetTopicSolutionUseCase(dto dto.ChooseSnippetTopicSolutionDTO) error
	GetSnippetTopicUseCase(dto dto.GetSnippetTopicInputDTO) (*dto.GetSnippetTopicOutputDTO, error)
	GetSnippetUseCase(dto dto.GetSnippetInputDTO) (*dto.GetSnippetOutputDTO, error)
	CompareSnippetVersionsModificationsUseCase(dto dto.CompareSnippetVersionsModificationsInputDTO) (*dto.CompareSnippetVersionsModificationsOutputDTO, error)
	ModifySnippetTopicStatusUseCase(dto dto.ModifySnippetTopicStatusInputDTO) error
	DeleteSnippetTopicUseCase(dto dto.DeleteSnippetTopicInputDTO) error

	// # -
	// COMMENTS AND VOTES
	// # -
	// ChooseSnippetTopicBestAnswerUseCase()
	// CommentOnSnippetUseCase()
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
	//snippetRepository repository.SnippetRepository,
	snippetTopicRepository repository.SnippetTopicRepository,
	accountClient adapter.AccountClient,
) *ReviewUseCaseRegistry {
	return &ReviewUseCaseRegistry{
		//SnippetRepository:      snippetRepository,
		SnippetTopicRepository: snippetTopicRepository,
		AccountClient:          accountClient,
	}
}

type ReviewUseCaseRegistry struct {
	//SnippetRepository      repository.SnippetRepository
	SnippetTopicRepository repository.SnippetTopicRepository
	AccountClient          adapter.AccountClient
}
