package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_VoteOnCommentUseCase() {
	s.Run("it should be able to vote on comment", func() {
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

		isUp := true

		input := dto.VoteOnCommentInputDTO{
			CommentID: comment.ID,
			AccountID: acc.ID,
			IsUp:      isUp,
		}

		err = s.useCase.VoteOnCommentUseCase(input)
		s.NoError(err)

		modifiedTopic := s.snippetTopicRepo.Items[0]
		s.Equal(topic.ID, modifiedTopic.ID)

		modifiedComment := s.commentRepo.Items[0]
		s.Equal(comment.ID, modifiedComment.ID)
		s.Equal(comment.Votes+1, modifiedComment.Votes)
	})

	s.Run("it should be not able to vote on comment if account doesn't exists", func() {
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

		isUp := true

		input := dto.VoteOnCommentInputDTO{
			CommentID: comment.ID,
			AccountID: invalidAccountID,
			IsUp:      isUp,
		}

		err = s.useCase.VoteOnCommentUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to vote on comment if account isn't valid", func() {
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

		isUp := true

		input := dto.VoteOnCommentInputDTO{
			CommentID: comment.ID,
			AccountID: acc.ID,
			IsUp:      isUp,
		}

		err = s.useCase.VoteOnCommentUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to vote on comment if comment doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		invalidCommentID := "invalid id"
		_, err = factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		isUp := true

		input := dto.VoteOnCommentInputDTO{
			CommentID: invalidCommentID,
			AccountID: acc.ID,
			IsUp:      isUp,
		}

		err = s.useCase.VoteOnCommentUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("comment").Error(), err.Error())
	})

	s.Run("it should be not able to vote on comment if account already voted in the comment", func() {
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

		isUp := true

		_, err = factory.MakeCommentVote(s.commentVoteRepo, factory.MakeCommentVoteInput{
			AccountID: acc.ID,
			CommentID: comment.ID,
			IsUp:      isUp,
		})
		s.NoError(err)

		input := dto.VoteOnCommentInputDTO{
			CommentID: comment.ID,
			AccountID: acc.ID,
			IsUp:      isUp,
		}

		err = s.useCase.VoteOnCommentUseCase(input)
		s.Error(err)
		s.Equal(core.NewAlreadyExistsErr("comment vote").Error(), err.Error())
	})
}
