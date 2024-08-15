package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_SubmitNewSnippetVersionUseCase() {
	codeSnippet := `
	export default function Home() {
		return(
			<div>
				<h1>hello world</div>
			</div>
		)
	}`
	message := "Renders successfully a home component"

	s.Run("it should be able to submit a new snippet version to a topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		input := dto.SubmitNewSnippetVersionInputDTO{
			Message:     message,
			CodeSnippet: codeSnippet,
			TopicID:     topic.ID,
			AccountID:   acc.ID,
		}

		err = s.useCase.SubmitNewSnippetVersionUseCase(input)

		s.NoError(err)

		snippet := s.snippetRepo.Items[0]

		s.Equal(codeSnippet, snippet.CodeSnippet)
		s.Equal(message, snippet.Message)
		s.Equal(topic.ID, snippet.TopicID)
	})

	s.Run("it should be not able to submit a new snippet version to a topic if account doesn't exists", func() {
		invalidAccountID := "invalid id"

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: invalidAccountID,
		})
		s.NoError(err)

		input := dto.SubmitNewSnippetVersionInputDTO{
			Message:     message,
			CodeSnippet: codeSnippet,
			TopicID:     topic.ID,
			AccountID:   invalidAccountID,
		}

		err = s.useCase.SubmitNewSnippetVersionUseCase(input)

		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to submit a new snippet version to a topic if account isn't valid", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  false,
			Verified: false,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		input := dto.SubmitNewSnippetVersionInputDTO{
			Message:     message,
			CodeSnippet: codeSnippet,
			TopicID:     topic.ID,
			AccountID:   acc.ID,
		}

		err = s.useCase.SubmitNewSnippetVersionUseCase(input)

		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to submit a new snippet version to a topic that don't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		input := dto.SubmitNewSnippetVersionInputDTO{
			Message:     message,
			CodeSnippet: codeSnippet,
			TopicID:     "invalid id",
			AccountID:   acc.ID,
		}

		err = s.useCase.SubmitNewSnippetVersionUseCase(input)

		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet topic").Error(), err.Error())
	})

	s.Run("it should be not able to submit a new snippet version to a topic if it's not the topic owner", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: "another account id",
		})
		s.NoError(err)

		input := dto.SubmitNewSnippetVersionInputDTO{
			Message:     message,
			CodeSnippet: codeSnippet,
			TopicID:     topic.ID,
			AccountID:   acc.ID,
		}

		err = s.useCase.SubmitNewSnippetVersionUseCase(input)

		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to submit a new snippet version to a topic if is with invalid payload", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		input := dto.SubmitNewSnippetVersionInputDTO{
			Message:     message,
			CodeSnippet: "",
			TopicID:     topic.ID,
			AccountID:   acc.ID,
		}

		err = s.useCase.SubmitNewSnippetVersionUseCase(input)

		s.Error(err)
		s.Equal(core.NewValidationErr(core.ErrRequired("codesnippet")).Error(), err.Error())
	})
}
