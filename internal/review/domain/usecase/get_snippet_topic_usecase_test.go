package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_GetSnippetTopicUseCase() {
	s.Run("it should be able to get snippet topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		input := dto.GetSnippetTopicInputDTO{SnippetTopicID: topic.ID}

		output, err := s.useCase.GetSnippetTopicUseCase(input)
		s.NoError(err)
		s.Equal(topic.ID, output.SnippetTopic.ID)
	})

	s.Run("it should be not able to get snippet topic if topic doesn't exists", func() {
		input := dto.GetSnippetTopicInputDTO{SnippetTopicID: "invalid id"}

		_, err := s.useCase.GetSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet topic").Error(), err.Error())
	})
}
