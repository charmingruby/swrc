package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_DeleteSnippetTopicUseCase() {
	s.Run("it should be able to delete a snippet topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)
		s.Equal(1, len(s.snippetTopicRepo.Items))

		input := dto.DeleteSnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		}

		err = s.useCase.DeleteSnippetTopicUseCase(input)
		s.NoError(err)
		s.Equal(0, len(s.snippetTopicRepo.Items))
	})

	s.Run("it should be able to delete a snippet topic even if have snippets attached", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)
		s.Equal(1, len(s.snippetTopicRepo.Items))

		_, err = factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		_, err = factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		s.Equal(2, len(s.snippetRepo.Items))

		input := dto.DeleteSnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		}

		err = s.useCase.DeleteSnippetTopicUseCase(input)
		s.NoError(err)
		s.Equal(0, len(s.snippetTopicRepo.Items))
		s.Equal(0, len(s.snippetRepo.Items))
	})

	s.Run("it should be not able to delete a snippet topic with an invalid account", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  false,
			Verified: false,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)
		s.Equal(1, len(s.snippetTopicRepo.Items))

		input := dto.DeleteSnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		}

		err = s.useCase.DeleteSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to delete a snippet topic if account doesn't exists", func() {
		invalidAccountID := "invalid id"

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: invalidAccountID,
		})
		s.NoError(err)
		s.Equal(1, len(s.snippetTopicRepo.Items))

		input := dto.DeleteSnippetTopicInputDTO{
			AccountID:      invalidAccountID,
			SnippetTopicID: topic.ID,
		}

		err = s.useCase.DeleteSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to delete a snippet topic if account_id is not the owner", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: "account_owner_id",
		})
		s.NoError(err)
		s.Equal(1, len(s.snippetTopicRepo.Items))

		input := dto.DeleteSnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		}

		err = s.useCase.DeleteSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to delete a snippet topic that doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		input := dto.DeleteSnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: "invalid topic id",
		}

		err = s.useCase.DeleteSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet topic").Error(), err.Error())
	})
}
