package usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_CreateSnippetTopicUseCase() {
	title := "snippet title"
	description := "snippet description"

	s.Run("it should be able to create a new snippet topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		input := dto.CreateSnippetTopicInputDTO{
			Title:       title,
			Description: description,
			AccountID:   acc.ID,
		}

		err = s.useCase.CreateSnippetTopicUseCase(input)
		s.NoError(err)

		s.Equal(title, s.snippetTopicRepo.Items[0].Title)
		s.Equal(description, s.snippetTopicRepo.Items[0].Description)
		s.Equal(acc.ID, s.snippetTopicRepo.Items[0].AccountID)
	})

	s.Run("it should be not able to create a snippet topic if account doesn't exists", func() {
		input := dto.CreateSnippetTopicInputDTO{
			Title:       title,
			Description: description,
			AccountID:   "invalid id",
		}

		err := s.useCase.CreateSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewNotFoundErr("account").Error(), err.Error())
	})

	s.Run("it should be not able to create a snippet topic if account isn't valid", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  false,
			Verified: false,
		})
		s.NoError(err)

		input := dto.CreateSnippetTopicInputDTO{
			Title:       title,
			Description: description,
			AccountID:   acc.ID,
		}

		err = s.useCase.CreateSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewUnauthorizedErr().Error(), err.Error())
	})

	s.Run("it should be not able to create a snippet topic with invalid params", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{
			IsValid:  true,
			Verified: true,
		})
		s.NoError(err)

		input := dto.CreateSnippetTopicInputDTO{
			Title:       "",
			Description: description,
			AccountID:   acc.ID,
		}

		err = s.useCase.CreateSnippetTopicUseCase(input)
		s.Error(err)
		s.Equal(core.NewValidationErr(core.ErrRequired("title")).Error(), err.Error())
	})
}
