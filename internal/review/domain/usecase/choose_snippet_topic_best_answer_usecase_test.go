package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_ChooseSnippetTopicBestAnswerUseCase() {
	s.Run("it should be able to choose a snippet topic best answer", func() {
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

		input := dto.ChooseSnippetTopicBestAnswerInputDTO{
			SnippetTopicID: topic.ID,
			CommentID:      comment.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicBestAnswerUseCase(input)
		s.NoError(err)

		modifiedTopic := s.snippetTopicRepo.Items[0]

		s.Equal(comment.ID, modifiedTopic.BestAnswerID)
	})

	s.Run("it should be not able to choose a snippet topic best answer if account doesn't exists", func() {
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

		input := dto.ChooseSnippetTopicBestAnswerInputDTO{
			SnippetTopicID: topic.ID,
			CommentID:      comment.ID,
			AccountID:      invalidAccountID,
		}

		err = s.useCase.ChooseSnippetTopicBestAnswerUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic best answer if account isn't valid", func() {
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

		input := dto.ChooseSnippetTopicBestAnswerInputDTO{
			SnippetTopicID: topic.ID,
			CommentID:      comment.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicBestAnswerUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic best answer if topic doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		invalidTopicID := "invalid id"

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc.ID,
			SnippetTopicID: invalidTopicID,
		})
		s.NoError(err)

		input := dto.ChooseSnippetTopicBestAnswerInputDTO{
			SnippetTopicID: invalidTopicID,
			CommentID:      comment.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicBestAnswerUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet topic").Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic best answer if account is not the topic owner", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: "not_owner_account id",
		})
		s.NoError(err)

		comment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.ChooseSnippetTopicBestAnswerInputDTO{
			SnippetTopicID: topic.ID,
			CommentID:      comment.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicBestAnswerUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic best answer if comment doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		invalidCommentID := "invalid id"

		input := dto.ChooseSnippetTopicBestAnswerInputDTO{
			SnippetTopicID: topic.ID,
			CommentID:      invalidCommentID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicBestAnswerUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("comment").Error(), err.Error())
	})

	s.Run("it should be not able to choose a snippet topic best answer if comment isn't from topic", func() {
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
			SnippetTopicID: "another topic",
		})
		s.NoError(err)

		input := dto.ChooseSnippetTopicBestAnswerInputDTO{
			SnippetTopicID: topic.ID,
			CommentID:      comment.ID,
			AccountID:      acc.ID,
		}

		err = s.useCase.ChooseSnippetTopicBestAnswerUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("comment").Error(), err.Error())
	})
}
