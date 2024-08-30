package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_RemoveVoteFromCommentUseCase() {
	s.Run("it should be able to remove a vote from a comment", func() {
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

		vote, err := factory.MakeCommentVote(s.commentVoteRepo, factory.MakeCommentVoteInput{
			IsUp:      true,
			AccountID: acc.ID,
			CommentID: comment.ID,
		})
		s.NoError(err)
		s.Equal(len(s.commentVoteRepo.Items), 1)
		s.Equal(vote.ID, s.commentVoteRepo.Items[0].ID)

		input := dto.RemoveVoteFromCommentInputDTO{
			AccountID:     acc.ID,
			CommentVoteID: vote.ID,
		}

		err = s.useCase.RemoveVoteFromCommentUseCase(input)
		s.NoError(err)
		s.Equal(len(s.commentVoteRepo.Items), 0)
	})

	s.Run("it should be not able to remove a vote from a comment if account isn't valid", func() {
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

		vote, err := factory.MakeCommentVote(s.commentVoteRepo, factory.MakeCommentVoteInput{
			IsUp:      true,
			AccountID: acc.ID,
			CommentID: comment.ID,
		})
		s.NoError(err)

		input := dto.RemoveVoteFromCommentInputDTO{
			AccountID:     acc.ID,
			CommentVoteID: vote.ID,
		}

		err = s.useCase.RemoveVoteFromCommentUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to remove a vote from a comment if account doesn't exists", func() {
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

		vote, err := factory.MakeCommentVote(s.commentVoteRepo, factory.MakeCommentVoteInput{
			IsUp:      true,
			AccountID: invalidAccountID,
			CommentID: comment.ID,
		})
		s.NoError(err)

		input := dto.RemoveVoteFromCommentInputDTO{
			AccountID:     invalidAccountID,
			CommentVoteID: vote.ID,
		}

		err = s.useCase.RemoveVoteFromCommentUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to remove a vote from a comment if vote doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		_, err = factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		invalidVoteID := "invalid id"

		input := dto.RemoveVoteFromCommentInputDTO{
			AccountID:     acc.ID,
			CommentVoteID: invalidVoteID,
		}

		err = s.useCase.RemoveVoteFromCommentUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("comment vote").Error(), err.Error())
	})

	s.Run("it should be not able to remove a vote from a comment if account isn't vote owner", func() {
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

		vote, err := factory.MakeCommentVote(s.commentVoteRepo, factory.MakeCommentVoteInput{
			IsUp:      true,
			AccountID: "owner_id",
			CommentID: comment.ID,
		})
		s.NoError(err)

		input := dto.RemoveVoteFromCommentInputDTO{
			AccountID:     acc.ID,
			CommentVoteID: vote.ID,
		}

		err = s.useCase.RemoveVoteFromCommentUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})
}
