package review_usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/review_dto"
)

type ReviewUseCase interface {
	// # -
	// SNIPPETS
	// # -
	CreateSnippetTopicUseCase(dto review_dto.CreateSnippetTopicInputDTO) (review_dto.CreateSnippetTopicOutputDTO, error)
	SubmitNewSnippetVersionUseCase(dto review_dto.SubmitNewSnippetVersionInputDTO) error
	ChooseSnippetTopicSolutionUseCase(dto review_dto.ChooseSnippetTopicSolutionDTO) error
	GetSnippetTopicUseCase(dto review_dto.GetSnippetTopicInputDTO) (review_dto.GetSnippetTopicOutputDTO, error)
	GetSnippetUseCase(dto review_dto.GetSnippetInputDTO) (review_dto.GetSnippetOutputDTO, error)
	CompareSnippetVersionsModificationsUseCase(dto review_dto.CompareSnippetVersionsModificationsInputDTO) (review_dto.CompareSnippetVersionsModificationsOutputDTO, error)
	ModifySnippetTopicStatusUseCase(dto review_dto.ModifySnippetTopicStatusInputDTO) error
	DeleteSnippetTopicUseCase(dto review_dto.DeleteSnippetTopicInputDTO) error

	// # -
	// COMMENTS AND VOTES
	// # -
	// ChooseSnippetTopicBestAnswerUseCase()
	// CommentOnSnippetUseCase()
	// VoteOnRootCommentUseCase()
	// RemoveVoteFromRootCommentUseCase()

	// # -
	// REACTIONS
	// # -
	// ReactOnSnippetUseCase()
	// RemoveReactionFromSnippetUseCase()

	// # -
	// LIST
	// # -
	// FetchOpenSnippetsUseCase()
	// FetchAccountSnippetsUseCase()
	// FetchClosedSnippetsUseCase()
}

func NewReviewUseCaseRegistry() {}

type ReviewUseCaseRegistry struct{}
