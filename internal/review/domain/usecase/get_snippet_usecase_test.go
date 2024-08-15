package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_GetSnippetUseCase() {
	s.Run("it should be able to get snippet", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		snippet, err := factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.GetSnippetInputDTO{SnippetID: snippet.ID}

		output, err := s.useCase.GetSnippetUseCase(input)
		s.NoError(err)
		s.Equal(snippet.ID, output.Snippet.ID)
		s.Equal(snippet.Version, output.Snippet.Version)
	})

	s.Run("it should be not able to get snippet if snippet doesn't exists", func() {
		input := dto.GetSnippetInputDTO{SnippetID: "invalid id"}

		_, err := s.useCase.GetSnippetUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet").Error(), err.Error())
	})
}
