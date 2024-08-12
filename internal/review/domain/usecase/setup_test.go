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
	reviewUseCase    *ReviewUseCaseRegistry
	snippetTopicRepo *inmemory_repository.InMemorySnippetTopicRepository
	accountRepo      *inmemory_repository.InMemoryAccountRepository
}

func (s *Suite) SetupSuite() {
	s.snippetTopicRepo = inmemory_repository.NewInMemorySnippetTopicRepository()
	s.accountRepo = inmemory_repository.NewInMemoryAccountRepository()

	accountClient := client.NewAccountClient(s.accountRepo)

	s.reviewUseCase = NewReviewUseCaseRegistry(s.snippetTopicRepo, accountClient)
}

func (s *Suite) SetupTest() {
	s.snippetTopicRepo.Items = []entity.SnippetTopic{}
	s.accountRepo.Items = []accountEntity.Account{}
}

func (s *Suite) TearDownTest() {
	s.snippetTopicRepo.Items = []entity.SnippetTopic{}
	s.accountRepo.Items = []accountEntity.Account{}
}

func (s *Suite) SetupSubTest() {
	s.snippetTopicRepo.Items = []entity.SnippetTopic{}
	s.accountRepo.Items = []accountEntity.Account{}
}

func (s *Suite) TearDownSubTest() {
	s.snippetTopicRepo.Items = []entity.SnippetTopic{}
	s.accountRepo.Items = []accountEntity.Account{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
