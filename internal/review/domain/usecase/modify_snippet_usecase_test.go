package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_ModifySnippetUseCase() {
	s.Run("it should be able to modify snippet", func() {
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

		newCodeSnippet := "new" + snippet.CodeSnippet
		newMessage := "new" + snippet.Message

		input := dto.ModifySnippetInputDTO{
			AccountID:   acc.ID,
			SnippetID:   snippet.ID,
			Message:     newMessage,
			CodeSnippet: newCodeSnippet,
		}

		err = s.useCase.ModifySnippetUseCase(input)
		s.NoError(err)

		snippetFound := s.snippetRepo.Items[0]
		s.Equal(newMessage, snippetFound.Message)
		s.Equal(newCodeSnippet, snippetFound.CodeSnippet)
	})

	s.Run("it should be not able to modify snippet if account isn't valid", func() {
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

		newCodeSnippet := "new" + snippet.CodeSnippet
		newMessage := "new" + snippet.Message

		input := dto.ModifySnippetInputDTO{
			AccountID:   acc.ID,
			SnippetID:   snippet.ID,
			Message:     newMessage,
			CodeSnippet: newCodeSnippet,
		}

		err = s.useCase.ModifySnippetUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to modify snippet if account doesn't exists", func() {
		invalidAccountID := "invalid id"

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: invalidAccountID,
		})
		s.NoError(err)

		snippet, err := factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		newCodeSnippet := "new" + snippet.CodeSnippet
		newMessage := "new" + snippet.Message

		input := dto.ModifySnippetInputDTO{
			AccountID:   invalidAccountID,
			SnippetID:   snippet.ID,
			Message:     newMessage,
			CodeSnippet: newCodeSnippet,
		}

		err = s.useCase.ModifySnippetUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to modify snippet if snippet topic doesn't exists", func() {
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

		newCodeSnippet := "new" + snippet.CodeSnippet
		newMessage := "new" + snippet.Message

		input := dto.ModifySnippetInputDTO{
			AccountID:   acc.ID,
			SnippetID:   snippet.ID,
			Message:     newMessage,
			CodeSnippet: newCodeSnippet,
		}

		err = s.useCase.ModifySnippetUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet topic").Error(), err.Error())
	})

	s.Run("it should be not able to modify snippet if is not the snippet owner", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: "invalid id",
		})
		s.NoError(err)

		snippet, err := factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		newCodeSnippet := "new" + snippet.CodeSnippet
		newMessage := "new" + snippet.Message

		input := dto.ModifySnippetInputDTO{
			AccountID:   acc.ID,
			SnippetID:   snippet.ID,
			Message:     newMessage,
			CodeSnippet: newCodeSnippet,
		}

		err = s.useCase.ModifySnippetUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to modify snippet if doesn't have anything to change", func() {
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

		input := dto.ModifySnippetInputDTO{
			AccountID:   acc.ID,
			SnippetID:   snippet.ID,
			Message:     snippet.Message,
			CodeSnippet: snippet.CodeSnippet,
		}

		err = s.useCase.ModifySnippetUseCase(input)
		s.Error(err)
		s.Equal(core.NewNothingToChangeErr().Error(), err.Error())
	})
}
