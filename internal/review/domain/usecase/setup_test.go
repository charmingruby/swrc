package usecase

import (
	"testing"

	accountEntity "github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/domain/client"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/test/inmemory_repository"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	useCase          *ReviewUseCaseRegistry
	snippetTopicRepo *inmemory_repository.InMemorySnippetTopicRepository
	snippetRepo      *inmemory_repository.InMemorySnippetRepository
	commentVoteRepo  *inmemory_repository.InMemoryCommentVoteRepository
	commentRepo      *inmemory_repository.InMemoryCommentRepository
	accountRepo      *inmemory_repository.InMemoryAccountRepository
}

func (s *Suite) SetupSuite() {
	s.snippetTopicRepo = inmemory_repository.NewInMemorySnippetTopicRepository()
	s.snippetRepo = inmemory_repository.NewInMemorySnippetRepository()
	s.commentRepo = inmemory_repository.NewInMemoryCommentRepository()
	s.accountRepo = inmemory_repository.NewInMemoryAccountRepository()
	s.commentVoteRepo = inmemory_repository.NewInMemoryCommentVoteRepository()

	accountClient := client.NewAccountClient(s.accountRepo)

	s.useCase = NewReviewUseCaseRegistry(s.snippetRepo, s.snippetTopicRepo, s.commentRepo, s.commentVoteRepo, accountClient)
}

func (s *Suite) SetupTest() {
	s.snippetTopicRepo.Items = []entity.SnippetTopic{}
	s.snippetRepo.Items = []entity.Snippet{}
	s.accountRepo.Items = []accountEntity.Account{}
	s.commentRepo.Items = []entity.Comment{}
	s.commentVoteRepo.Items = []entity.CommentVote{}
}

func (s *Suite) TearDownTest() {
	s.snippetTopicRepo.Items = []entity.SnippetTopic{}
	s.snippetRepo.Items = []entity.Snippet{}
	s.accountRepo.Items = []accountEntity.Account{}
	s.commentRepo.Items = []entity.Comment{}
	s.commentVoteRepo.Items = []entity.CommentVote{}
}

func (s *Suite) SetupSubTest() {
	s.snippetTopicRepo.Items = []entity.SnippetTopic{}
	s.snippetRepo.Items = []entity.Snippet{}
	s.accountRepo.Items = []accountEntity.Account{}
	s.commentRepo.Items = []entity.Comment{}
	s.commentVoteRepo.Items = []entity.CommentVote{}
}

func (s *Suite) TearDownSubTest() {
	s.snippetTopicRepo.Items = []entity.SnippetTopic{}
	s.snippetRepo.Items = []entity.Snippet{}
	s.accountRepo.Items = []accountEntity.Account{}
	s.commentRepo.Items = []entity.Comment{}
	s.commentVoteRepo.Items = []entity.CommentVote{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
