package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_ChooseSnippetTopicSolutionUseCase() {
	s.Run("it should be able to choose a snippet topic solution", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		snippet, err := factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.ChooseSnippetTopicSolutionInputDTO{
			SnippetTopicID: topic.ID,
			SnippetID:      snippet.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicSolutionUseCase(input)
		s.NoError(err)

		modifiedTopic := s.snippetTopicRepo.Items[0]

		s.Equal(snippet.ID, modifiedTopic.SnippetSolutionID)
	})

	s.Run("it should be not able to choose a snippet topic solution if account doesn't exists", func() {
		invalidAccountID := "invalid id"

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: invalidAccountID,
		})
		s.NoError(err)

		snippet, err := factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.ChooseSnippetTopicSolutionInputDTO{
			SnippetTopicID: topic.ID,
			SnippetID:      snippet.ID,
			AccountID:      invalidAccountID,
		}

		err = s.useCase.ChooseSnippetTopicSolutionUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic solution if account isn't valid", func() {
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

		input := dto.ChooseSnippetTopicSolutionInputDTO{
			SnippetTopicID: topic.ID,
			SnippetID:      snippet.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicSolutionUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic solution if topic doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		invalidTopicID := "invalid id"

		snippet, err := factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: invalidTopicID,
		})
		s.NoError(err)

		input := dto.ChooseSnippetTopicSolutionInputDTO{
			SnippetTopicID: invalidTopicID,
			SnippetID:      snippet.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicSolutionUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet topic").Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic solution if account is not the topic owner", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: "not_owner_account id",
		})
		s.NoError(err)

		snippet, err := factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.ChooseSnippetTopicSolutionInputDTO{
			SnippetTopicID: topic.ID,
			SnippetID:      snippet.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicSolutionUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic solution if snippet doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		invalidSnippetID := "invalid id"

		input := dto.ChooseSnippetTopicSolutionInputDTO{
			SnippetTopicID: topic.ID,
			SnippetID:      invalidSnippetID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicSolutionUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet").Error(), err.Error())
	})
}
