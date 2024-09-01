package entity

import (
	"testing"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewSnippetTopic(t *testing.T) {
	title := "snippet title"
	description := "snippet description"
	accountID := "account id"

	t.Run("it should be able to create a new snippet topic", func(t *testing.T) {
		snp, err := NewSnippetTopic(title, description, accountID)

		assert.NoError(t, err)
		assert.Equal(t, title, snp.Title)
		assert.Equal(t, description, snp.Description)
		assert.Equal(t, 0, snp.CurrentVersion)
		assert.Equal(t, "", snp.BestAnswerID)
		assert.Equal(t, SNIPPET_TOPIC_STATUS_OPEN, snp.Status)
		assert.Equal(t, accountID, snp.AccountID)
	})

	t.Run("it should be not able to create a snippet topic with blank title", func(t *testing.T) {
		snp, err := NewSnippetTopic("", description, accountID)

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("title")).Error(), err.Error())
	})

	t.Run("it should be notf able to create a snippet topic with blank description", func(t *testing.T) {
		snp, err := NewSnippetTopic(title, "", accountID)

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("description")).Error(), err.Error())
	})

	t.Run("it should be not able to create a snippet topic with blank account id", func(t *testing.T) {
		snp, err := NewSnippetTopic(title, description, "")

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("accountid")).Error(), err.Error())
	})
}

func Test_SnippetTopicClose(t *testing.T) {
	t.Run("it should be able to close a snippet topic", func(t *testing.T) {
		snp, err := NewSnippetTopic("title", "description", "account_id")

		assert.NoError(t, err)
		assert.Equal(t, SNIPPET_TOPIC_STATUS_OPEN, snp.Status)

		snp.Close()

		assert.Equal(t, SNIPPET_TOPIC_STATUS_CLOSED, snp.Status)
	})
}
