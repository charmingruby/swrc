package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_CommentOnSnippetTopicUseCase() {
	content := "dummy content"

	s.Run("it should be able to comment on a snippet topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		noParentCommentID := ""

		input := dto.CommentOnSnippetTopicInputDTO{
			Content:         content,
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: noParentCommentID,
		}

		err = s.useCase.CommentOnSnippetTopicUseCase(input)
		s.NoError(err)

		commentFound := s.commentRepo.Items[0]
		s.Equal(content, commentFound.Content)
		s.Equal(noParentCommentID, commentFound.ParentCommentID)
		s.Equal(acc.ID, commentFound.AccountID)
	})

	s.Run("it should be able to answer a comment of a snippet topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		rootComment, err := factory.MakeComment(s.commentRepo, factory.MakeCommentInput{
			AccountID:      acc.ID,
			SnippetTopicID: topic.ID,
		})
		s.NoError(err)

		input := dto.CommentOnSnippetTopicInputDTO{
			Content:         content,
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: rootComment.ID,
		}

		err = s.useCase.CommentOnSnippetTopicUseCase(input)
		s.NoError(err)

		commentFound := s.commentRepo.Items[1]
		s.Equal(content, commentFound.Content)
		s.Equal(rootComment.ID, commentFound.ParentCommentID)
		s.Equal(acc.ID, commentFound.AccountID)
	})

	s.Run("it should be not able to comment on a snippet topic if account isn't valid", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		noParentCommentID := ""

		input := dto.CommentOnSnippetTopicInputDTO{
			Content:         content,
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: noParentCommentID,
		}

		err = s.useCase.CommentOnSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to comment on a snippet topic if account doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		invalidAccountID := "invalid id"
		noParentCommentID := ""

		input := dto.CommentOnSnippetTopicInputDTO{
			Content:         content,
			AccountID:       invalidAccountID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: noParentCommentID,
		}

		err = s.useCase.CommentOnSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to comment on a snippet topic if topic doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		invalidTopicID := "invalid id"

		noParentCommentID := ""

		input := dto.CommentOnSnippetTopicInputDTO{
			Content:         content,
			AccountID:       acc.ID,
			SnippetTopicID:  invalidTopicID,
			ParentCommentID: noParentCommentID,
		}

		err = s.useCase.CommentOnSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("snippet topic").Error(), err.Error())
	})

	s.Run("it should be not able to answer comment of a snippet topic if comment doesn't exists", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		invalidParentCommentID := "invalid id"

		input := dto.CommentOnSnippetTopicInputDTO{
			Content:         content,
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: invalidParentCommentID,
		}

		err = s.useCase.CommentOnSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("parent comment").Error(), err.Error())
	})

	s.Run("it should be not able to comment on a snippet topic if comment params is invalid", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		topic, err := factory.MakeSnippetTopic(s.snippetTopicRepo, factory.MakeSnippetTopicInput{
			AccountID: acc.ID,
		})
		s.NoError(err)

		noParentCommentID := ""

		input := dto.CommentOnSnippetTopicInputDTO{
			Content:         "",
			AccountID:       acc.ID,
			SnippetTopicID:  topic.ID,
			ParentCommentID: noParentCommentID,
		}

		err = s.useCase.CommentOnSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewValidationErr(core.ErrRequired("content")).Error(), err.Error())
	})
}
