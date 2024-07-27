package review_entity

import (
	"testing"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewSnippet(t *testing.T) {
	title := "snippet title"
	description := "snippet description"
	codeSnippet := "code snippet"
	accountID := "account id"

	t.Run("it should be able to create a new snippet", func(t *testing.T) {
		snp, err := NewSnippet(title, description, codeSnippet, accountID)

		assert.Equal(t, title, snp.Title)
		assert.Equal(t, description, snp.Description)
		assert.Equal(t, codeSnippet, snp.CodeSnippet)
		assert.Equal(t, SNIPPET_STATUS_OPEN, snp.Status)
		assert.NoError(t, err)
		assert.Equal(t, 0, snp.Reactions)
		assert.Equal(t, 0, snp.Comments)
		assert.Equal(t, 0, snp.Votes)
		assert.Equal(t, accountID, snp.AccountID)
	})

	t.Run("it should be not able to create a snippet with blank title", func(t *testing.T) {
		snp, err := NewSnippet("", description, codeSnippet, accountID)

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("title")).Error(), err.Error())
	})

	t.Run("it should be notf able to create a snippet with blank description", func(t *testing.T) {
		snp, err := NewSnippet(title, "", codeSnippet, accountID)

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("description")).Error(), err.Error())
	})

	t.Run("it should be not able to create a snippet with blank code snippet", func(t *testing.T) {
		snp, err := NewSnippet(title, description, "", accountID)

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("codesnippet")).Error(), err.Error())
	})

	t.Run("it should be not able to create a snippet with blank account id", func(t *testing.T) {
		snp, err := NewSnippet(title, description, codeSnippet, "")

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("accountid")).Error(), err.Error())
	})
}

func Test_SnippetClose(t *testing.T) {
	t.Run("it should be able to close a snippet", func(t *testing.T) {
		snp, err := NewSnippet("title", "description", "print(`hello world`)", "account_id")

		assert.NoError(t, err)
		assert.Equal(t, SNIPPET_STATUS_OPEN, snp.Status)

		snp.Close()

		assert.Equal(t, SNIPPET_STATUS_CLOSED, snp.Status)
	})
}
