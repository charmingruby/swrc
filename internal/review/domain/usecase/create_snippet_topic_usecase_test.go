package usecase

import (
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/test/factory"
)

func (s *Suite) Test_CreateSnippetTopicUseCase() {
	title := "snippet title"
	description := "snippet description"

	s.Run("it should be able to create a new snippet topic", func() {
		acc, err := factory.MakeAccount(s.accountRepo, factory.MakeAccountInput{})
		s.NoError(err)

		input := dto.CreateSnippetTopicInputDTO{
			Title:       title,
			Description: description,
			AccountID:   acc.ID,
		}

		err = s.reviewUseCase.CreateSnippetTopicUseCase(input)
		s.NoError(err)
	})

	s.Run("it should be not able to create a snippet topic if account doesn't exists", func() {

	})

	s.Run("it should be not able to create a snippet topic if account isn't valid", func() {

	})

	s.Run("it should be not able to create a snippet topic with invalid params", func() {

	})
}
