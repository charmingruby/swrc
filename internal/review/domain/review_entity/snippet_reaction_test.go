package review_entity

import (
	"testing"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewSnippetReaction(t *testing.T) {
	snippetID := "snippet id"
	accountID := "account id"
	defaultReaction := SNIPPET_REACTION_LIKE

	t.Run("it should be able to create a new snippet like reaction", func(t *testing.T) {
		r, err := NewSnippetReaction("LIKE", snippetID, accountID)

		assert.NoError(t, err)
		assert.Equal(t, SNIPPET_REACTION_LIKE, r.Reaction)
		assert.Equal(t, accountID, r.AccountID)
		assert.Equal(t, snippetID, r.SnippetID)
	})

	t.Run("it should be able to create a new snippet challenging reaction", func(t *testing.T) {
		r, err := NewSnippetReaction("CHALLENGING", snippetID, accountID)

		assert.NoError(t, err)
		assert.Equal(t, SNIPPET_REACTION_CHALLENGING, r.Reaction)
		assert.Equal(t, accountID, r.AccountID)
		assert.Equal(t, snippetID, r.SnippetID)
	})

	t.Run("it should be able to create a new snippet mind blow reaction", func(t *testing.T) {
		r, err := NewSnippetReaction("MIND BLOW", snippetID, accountID)

		assert.NoError(t, err)
		assert.Equal(t, SNIPPET_REACTION_MINDBLOW, r.Reaction)
		assert.Equal(t, accountID, r.AccountID)
		assert.Equal(t, snippetID, r.SnippetID)
	})

	t.Run("it should be able to create a new snippet bug reaction", func(t *testing.T) {
		r, err := NewSnippetReaction("BUG", snippetID, accountID)

		assert.NoError(t, err)
		assert.Equal(t, SNIPPET_REACTION_BUG, r.Reaction)
		assert.Equal(t, accountID, r.AccountID)
		assert.Equal(t, snippetID, r.SnippetID)
	})

	t.Run("it should be not able to create a snippet reaction with an invalid snippet reaction", func(t *testing.T) {
		r, err := NewSnippetReaction("TEST", snippetID, accountID)

		assert.Error(t, err)
		assert.Nil(t, r)
		assert.Equal(t, core.NewValidationErr("invalid reaction").Error(), err.Error())
	})

	t.Run("it should be not able to create a snippet reaction with a blank reaction", func(t *testing.T) {
		r, err := NewSnippetReaction("", snippetID, accountID)

		assert.Error(t, err)
		assert.Nil(t, r)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("reaction")).Error(), err.Error())
	})

	t.Run("it should be not able to create a snippet reaction with a blank snippet id", func(t *testing.T) {
		r, err := NewSnippetReaction(defaultReaction, "", accountID)

		assert.Error(t, err)
		assert.Nil(t, r)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("snippetid")).Error(), err.Error())
	})

	t.Run("it should be not able to create a snippet reaction with a blank account id", func(t *testing.T) {
		r, err := NewSnippetReaction(defaultReaction, snippetID, "")

		assert.Error(t, err)
		assert.Nil(t, r)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("accountid")).Error(), err.Error())
	})
}
