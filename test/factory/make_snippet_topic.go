package factory

import (
	"github.com/charmingruby/swrc/internal/common/util"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/internal/review/domain/repository"
)

type MakeSnippetTopicInput struct {
	Title       string
	Description string
	AccountID   string
}

func MakeSnippetTopic(
	repo repository.SnippetTopicRepository,
	in MakeSnippetTopicInput,
) (entity.SnippetTopic, error) {
	title := util.Ternary[string](in.Title == "", "dummy title", in.Title)
	description := util.Ternary[string](in.Description == "", "dummy description", in.Description)
	accountID := util.Ternary[string](in.AccountID == "", "invalid id", in.AccountID)

	snippetTopic, err := entity.NewSnippetTopic(title, description, accountID)
	if err != nil {
		return entity.SnippetTopic{}, err
	}

	if err := repo.Store(*snippetTopic); err != nil {
		return entity.SnippetTopic{}, err
	}

	return *snippetTopic, nil
}
