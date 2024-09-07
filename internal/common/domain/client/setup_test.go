package client

import (
	"testing"

	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/review/domain/port"
	"github.com/charmingruby/swrc/test/inmemory_repository"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	accountRepository *inmemory_repository.InMemoryAccountRepository
	accountClient     port.AccountClient
}

func (s *Suite) SetupSuite() {
	s.accountRepository = inmemory_repository.NewInMemoryAccountRepository()
	s.accountClient = NewAccountClient(s.accountRepository)
}

func (s *Suite) SetupTest() {
	s.accountRepository.Items = []entity.Account{}
}

func (s *Suite) TearDownTest() {
	s.accountRepository.Items = []entity.Account{}
}

func (s *Suite) SetupSubTest() {
	s.accountRepository.Items = []entity.Account{}
}

func (s *Suite) TearDownSubTest() {
	s.accountRepository.Items = []entity.Account{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
