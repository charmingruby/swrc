package entity

import (
	"testing"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewComment(t *testing.T) {
	content := "dummy content"
	accountID := "account_id"
	parentCommentID := "parent_comment_id"
	snippetID := "snippet_id"

	t.Run("it should be able to create a new comment", func(t *testing.T) {
		c, err := NewComment(content, accountID, snippetID, parentCommentID)

		assert.NoError(t, err)
		assert.NotNil(t, c)
		assert.Equal(t, content, c.Content)
		assert.Equal(t, accountID, c.AccountID)
		assert.Equal(t, parentCommentID, c.ParentCommentID)
		assert.Equal(t, snippetID, c.SnippetTopicID)
		assert.Equal(t, 0, c.Votes)
	})

	t.Run("it should be able to create a new comment withou a parent comment id", func(t *testing.T) {
		c, err := NewComment(content, accountID, snippetID, "")

		assert.NoError(t, err)
		assert.NotNil(t, c)
		assert.Equal(t, content, c.Content)
		assert.Equal(t, accountID, c.AccountID)
		assert.Equal(t, "", c.ParentCommentID)
		assert.Equal(t, snippetID, c.SnippetTopicID)
		assert.Equal(t, 0, c.Votes)
	})

	t.Run("it should be not able to create a comment with blank content", func(t *testing.T) {
		c, err := NewComment("", accountID, snippetID, parentCommentID)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("content")).Error(), err.Error())
	})

	t.Run("it should be not able to create a comment with blank account id", func(t *testing.T) {
		c, err := NewComment(content, "", snippetID, parentCommentID)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("accountid")).Error(), err.Error())
	})

	t.Run("it should be not able to create a comment with blank snippet id", func(t *testing.T) {
		c, err := NewComment(content, accountID, "", parentCommentID)

		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("snippettopicid")).Error(), err.Error())
	})
}
