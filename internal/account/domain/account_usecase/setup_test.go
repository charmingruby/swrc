package account_usecase

import (
	"testing"

	"github.com/charmingruby/swrc/internal/account/domain/account_entity"
	"github.com/charmingruby/swrc/test/fake"
	"github.com/charmingruby/swrc/test/inmemory_repository"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	accountRepository *inmemory_repository.InMemoryAccountRepository
	fakeHashService   *fake.FakeHashService
	accountUseCase    *AccountUseCaseRegistry
}

func (s *Suite) SetupSuite() {
	s.accountRepository = inmemory_repository.NewInMemoryAccountRepository()
	s.fakeHashService = fake.NewFakeHashService()
	s.accountUseCase = NewAccountUseCaseRegistry(s.accountRepository, s.fakeHashService)
}

func (s *Suite) SetupTest() {
	s.accountRepository.Items = []account_entity.Account{}
}

func (s *Suite) TearDownTest() {
	s.accountRepository.Items = []account_entity.Account{}
}

func (s *Suite) SetupSubTest() {
	s.accountRepository.Items = []account_entity.Account{}
}

func (s *Suite) TearDownSubTest() {
	s.accountRepository.Items = []account_entity.Account{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
