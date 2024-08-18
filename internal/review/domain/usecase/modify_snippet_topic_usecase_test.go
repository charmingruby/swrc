package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_ModifySnippetTopicUseCase() {
	s.Run("it should be able to modify a snippet topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		newTitle := "new" + topic.Title
		newDescription := "new" + topic.Description
		newStatus := entity.SNIPPET_TOPIC_STATUS_CLOSED

		input := dto.ModifySnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
			Title:          newTitle,
			Description:    newDescription,
			Status:         newStatus,
		}

		err = s.useCase.ModifySnippetTopicUseCase(input)
		s.NoError(err)

		topicFound := s.snippetTopicRepo.Items[0]

		s.Equal(newTitle, topicFound.Title)
		s.Equal(newDescription, topicFound.Description)
		s.Equal(newStatus, topicFound.Status)
	})

	s.Run("it should be not able to modify a snippet topic if account is invalid", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		newTitle := "new" + topic.Title
		newDescription := "new" + topic.Description
		newStatus := entity.SNIPPET_TOPIC_STATUS_CLOSED

		input := dto.ModifySnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
			Title:          newTitle,
			Description:    newDescription,
			Status:         newStatus,
		}

		err = s.useCase.ModifySnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to modify a snippet topic if account doesn't exists", func() {
		invalidAccountID := "invalid id"

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: invalidAccountID,
		})
		s.NoError(err)

		newTitle := "new" + topic.Title
		newDescription := "new" + topic.Description
		newStatus := entity.SNIPPET_TOPIC_STATUS_CLOSED

		input := dto.ModifySnippetTopicInputDTO{
			AccountID:      invalidAccountID,
			SnippetTopicID: topic.ID,
			Title:          newTitle,
			Description:    newDescription,
			Status:         newStatus,
		}

		err = s.useCase.ModifySnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to modify a snippet topic if topic doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		input := dto.ModifySnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: "invalid id",
			Title:          "invalid topic title",
			Description:    "invalid topic description",
			Status:         entity.SNIPPET_TOPIC_STATUS_CLOSED,
		}

		err = s.useCase.ModifySnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet topic").Error(), err.Error())
	})

	s.Run("it should be not able to modify a snippet topic if account_id is not the topic owner", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: "owner_account_id",
		})
		s.NoError(err)

		newTitle := "new" + topic.Title
		newDescription := "new" + topic.Description
		newStatus := entity.SNIPPET_TOPIC_STATUS_CLOSED

		input := dto.ModifySnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
			Title:          newTitle,
			Description:    newDescription,
			Status:         newStatus,
		}

		err = s.useCase.ModifySnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to modify a snippet topic if has nothing to change", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		input := dto.ModifySnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
			Title:          topic.Title,
			Description:    topic.Description,
			Status:         topic.Status,
		}

		err = s.useCase.ModifySnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNothingToChangeErr().Error(), err.Error())
	})

	s.Run("it should be not able to modify a snippet topic if status is invalid", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		input := dto.ModifySnippetTopicInputDTO{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
			Title:          topic.Title,
			Description:    topic.Description,
			Status:         "invalid id",
		}

		err = s.useCase.ModifySnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewValidationErr("invalid status").Error(), err.Error())
	})
}
