package review

import (
	"github.com/charmingruby/swrc/internal/review/domain/port"
	"github.com/charmingruby/swrc/internal/review/domain/repository"
	"github.com/charmingruby/swrc/internal/review/domain/usecase"
	"github.com/charmingruby/swrc/internal/review/infra/transport/grpc/server"
	"google.golang.org/grpc"
)

func NewService(
	snippetRepository repository.SnippetRepository,
	snippetTopicRepository repository.SnippetTopicRepository,
	commentRepository repository.CommentRepository,
	commentVoteRepository repository.CommentVoteRepository,
	accountClient port.AccountClient,
) usecase.ReviewUseCase {
	return usecase.NewReviewUseCaseRegistry(
		snippetRepository,
		snippetTopicRepository,
		commentRepository,
		commentVoteRepository,
		accountClient,
	)
}

func NewGRPCHandler(srv *grpc.Server, reviewSvc usecase.ReviewUseCase) {
	server.NewReviewGRPCServerHandler(reviewSvc, srv).Register()
}
