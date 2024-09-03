package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_FetchCommentsUseCase() {
	s.Run("it should be able to fetch comments by topic id and account id", func() {
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
			Status:    entity.SNIPPET_TOPIC_STATUS_OPEN,
		})
		s.NoError(err)

		topic2, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_CLOSED,
		})
		s.NoError(err)

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc1.ID,
			SnippetTopicID: topic1.ID,
		})
		s.NoError(err)

		_, err = factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc1.ID,
			SnippetTopicID: topic2.ID,
		})
		s.NoError(err)

		_, err = factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc2.ID,
			SnippetTopicID: topic2.ID,
		})
		s.NoError(err)

		comments, err := s.useCase.FetchCommentsUseCase(dto.FetchCommentsInputDTO{
			AccountID:      acc1.ID,
			SnippetTopicID: topic1.ID,
		})
		s.NoError(err)
		s.Equal(1, len(comments.Comments))
		s.Equal(comment.ID, comments.Comments[0].ID)
	})

	s.Run("it should be able to fetch comments by id", func() {
		acc1, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_OPEN,
		})
		s.NoError(err)

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc1.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		_, err = factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc1.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		comments, err := s.useCase.FetchCommentsUseCase(dto.FetchCommentsInputDTO{
			ID: comment.ID,
		})
		s.NoError(err)
		s.Equal(1, len(comments.Comments))
		s.Equal(comment.ID, comments.Comments[0].ID)
	})

	s.Run("it should be able to fetch comments by parent id", func() {
		acc1, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc1.ID,
			Status:    entity.SNIPPET_TOPIC_STATUS_OPEN,
		})
		s.NoError(err)

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc1.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		childComment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:       acc1.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: comment.ID,
		})
		s.NoError(err)

		_, err = factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc1.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		comments, err := s.useCase.FetchCommentsUseCase(dto.FetchCommentsInputDTO{
			ParentCommentID: comment.ID,
		})
		s.NoError(err)
		s.Equal(1, len(comments.Comments))
		s.Equal(childComment.ID, comments.Comments[0].ID)
	})

	s.Run("it should be able to fetch comments even if the result is empty", func() {
		comments, err := s.useCase.FetchCommentsUseCase(dto.FetchCommentsInputDTO{
			ID: "invalid id",
		})
		s.NoError(err)
		s.Equal(0, len(comments.Comments))
	})
}
