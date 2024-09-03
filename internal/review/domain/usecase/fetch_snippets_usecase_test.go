package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_FetchSnippetsUseCase() {
	s.Run("it should be able to fetch snippets", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		snp, err := factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		_, err = factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: "not same id",
		})
		s.NoError(err)

		snps, err := s.useCase.FetchSnippetsUseCase(dto.FetchSnippetInputDTO{
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)
		s.Equal(1, len(snps.Snippets))
		s.Equal(snp.ID, snps.Snippets[0].ID)
	})

	s.Run("it should be able to fetch snippets even if result is empty", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		_, err = factory.MakeSnippet(s.snippetRepo, factory.MakeSnippetInput{
			TopicID: topic.ID,
		})
		s.NoError(err)

		snps, err := s.useCase.FetchSnippetsUseCase(dto.FetchSnippetInputDTO{
			SnippetTopicID: "invalid id",
		})
		s.NoError(err)
		s.Equal(0, len(snps.Snippets))
	})
}
