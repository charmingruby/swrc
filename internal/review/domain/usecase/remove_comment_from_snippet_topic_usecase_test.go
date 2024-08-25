package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_RemoveCommentFromSnippetTopicUseCase() {
	s.Run("it should be able to remove a comment from a snippet topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)
		s.Equal(1, len(s.commentRepo.Items))

		input := dto.RemoveCommentFromSnippetTopicInputDTO{
			AccountID: acc.ID,
			CommentID: comment.ID,
		}

		err = s.useCase.RemoveCommentFromSnippetTopicUseCase(input)
		s.NoError(err)
		s.Equal(0, len(s.commentRepo.Items))
	})

	s.Run("it should be able to remove a comment from a snippet topic and delete all children comments", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		parentComment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		_, err = factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: parentComment.ID,
		})
		s.NoError(err)

		_, err = factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: parentComment.ID,
		})
		s.NoError(err)

		s.Equal(3, len(s.commentRepo.Items))

		input := dto.RemoveCommentFromSnippetTopicInputDTO{
			AccountID: acc.ID,
			CommentID: parentComment.ID,
		}

		err = s.useCase.RemoveCommentFromSnippetTopicUseCase(input)
		s.NoError(err)
		s.Equal(0, len(s.commentRepo.Items))
	})

	s.Run("it should be able to remove a comment from a snippet topic and delete all children and nested comments", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		parentComment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		childrenComment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: parentComment.ID,
		})
		s.NoError(err)

		_, err = factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: childrenComment.ID,
		})
		s.NoError(err)

		s.Equal(3, len(s.commentRepo.Items))

		input := dto.RemoveCommentFromSnippetTopicInputDTO{
			AccountID: acc.ID,
			CommentID: parentComment.ID,
		}

		err = s.useCase.RemoveCommentFromSnippetTopicUseCase(input)
		s.NoError(err)
		s.Equal(0, len(s.commentRepo.Items))
	})

	s.Run("it should be not able to remove a comment from a snippet topic if account is invalid", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.RemoveCommentFromSnippetTopicInputDTO{
			AccountID: acc.ID,
			CommentID: comment.ID,
		}

		err = s.useCase.RemoveCommentFromSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to remove a comment from a snippet topic if account doesn't exists", func() {
		invalidAccountID := "invalid id"

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: invalidAccountID,
		})
		s.NoError(err)

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      invalidAccountID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.RemoveCommentFromSnippetTopicInputDTO{
			AccountID: invalidAccountID,
			CommentID: comment.ID,
		}

		err = s.useCase.RemoveCommentFromSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to remove a comment from a snippet topic if comment doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		_, err = factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		invalidCommentID := "invalid id"

		input := dto.RemoveCommentFromSnippetTopicInputDTO{
			AccountID: acc.ID,
			CommentID: invalidCommentID,
		}

		err = s.useCase.RemoveCommentFromSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("comment").Error(), err.Error())
	})

	s.Run("it should be not able to remove a comment from a snippet topic if account isn't the comment owner", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      "invalid id",
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.RemoveCommentFromSnippetTopicInputDTO{
			AccountID: acc.ID,
			CommentID: comment.ID,
		}

		err = s.useCase.RemoveCommentFromSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})
}
