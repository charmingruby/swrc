package factory

import (
	"github.com/charmingruby/swrc/internal/common/util"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/internal/review/domain/repository"
)

type MakeSnippetInput struct {
	Version     int
	CodeSnippet string
	Message     string
	TopicID     string
}

func MakeSnippet(
	repo repository.SnippetRepository,
	in MakeSnippetInput,
) (entity.Snippet, error) {
	version := util.Ternary[int](in.Version == 0, -1, in.Version)
	codeSnippet := util.Ternary[string](in.CodeSnippet == "", `
	export default function Home() {
		return(
			<div>
				<h1>hello world</div>
			</div>
		)
	}`, in.CodeSnippet)
	message := util.Ternary[string](in.Message == "", "Renders successfully a home component", in.Message)
	topicID := util.Ternary[string](in.TopicID == "", "invalid id", in.TopicID)

	snippet, err := entity.NewSnippet(version, codeSnippet, message, topicID)
	if err != nil {
		return entity.Snippet{}, err
	}

	if err := repo.Store(*snippet); err != nil {
		return entity.Snippet{}, err
	}

	return *snippet, nil
}
