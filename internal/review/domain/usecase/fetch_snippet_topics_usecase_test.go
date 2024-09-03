package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_FetchSnippetTopicsUseCase() {
	s.Run("it should be able to fetch snippet topics with all not unique params", func() {
		acc1, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		acc2, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		_, err = factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_OPEN,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_CLOSED,
		})
		s.NoError(err)

		_, err = factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc2.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_CLOSED,
		})
		s.NoError(err)

		topics, err := s.useCase.FetchSnippetTopicsUseCase(dto.FetchSnippetTopicsInputDTO{
			AccountID: acc1.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_CLOSED,
		})
		s.NoError(err)
		s.Equal(1, len(topics.Topics))
		s.Equal(topic.ID, topics.Topics[0].ID)
	})

	s.Run("it should be able to fetch snippet topics only with id", func() {
		acc1, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		acc2, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		_, err = factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
		})
		s.NoError(err)

		_, err = factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc2.ID,
		})
		s.NoError(err)

		topics, err := s.useCase.FetchSnippetTopicsUseCase(dto.FetchSnippetTopicsInputDTO{
			ID: topic.ID,
		})
		s.NoError(err)
		s.Equal(1, len(topics.Topics))
		s.Equal(topic.ID, topics.Topics[0].ID)
	})

	s.Run("it should be able to fetch snippet topics only with account_id", func() {
		acc1, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		acc2, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic1, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
		})
		s.NoError(err)

		topic2, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
		})
		s.NoError(err)

		_, err = factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc2.ID,
		})
		s.NoError(err)

		topics, err := s.useCase.FetchSnippetTopicsUseCase(dto.FetchSnippetTopicsInputDTO{
			AccountID: acc1.ID,
		})
		s.NoError(err)
		s.Equal(2, len(topics.Topics))
		s.Equal(topic1.ID, topics.Topics[0].ID)
		s.Equal(topic2.ID, topics.Topics[1].ID)
	})

	s.Run("it should be able to fetch snippet topics only with status", func() {
		acc1, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		acc2, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		_, err = factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_OPEN,
		})
		s.NoError(err)

		topic1, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_CLOSED,
		})
		s.NoError(err)

		topic2, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc2.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_CLOSED,
		})
		s.NoError(err)

		topics, err := s.useCase.FetchSnippetTopicsUseCase(dto.FetchSnippetTopicsInputDTO{
			Status: entity.SNIPPET_TOPIC_STATUS_CLOSED,
		})
		s.NoError(err)
		s.Equal(2, len(topics.Topics))
		s.Equal(topic1.ID, topics.Topics[0].ID)
		s.Equal(topic2.ID, topics.Topics[1].ID)
	})

	s.Run("it should be able to fetch snippet topics even if not match params", func() {
		topics, err := s.useCase.FetchSnippetTopicsUseCase(dto.FetchSnippetTopicsInputDTO{
			ID: "invalid id",
		})
		s.NoError(err)
		s.Equal(0, len(topics.Topics))
	})
}
