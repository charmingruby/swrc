package review_entity

import (
	"testing"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewSnippet(t *testing.T) {
	version := 1
	codeSnippet := "var dummy string"
	message := "dummy message"
	topicID := "topic id"

	t.Run("it should be able to create a new snippet", func(t *testing.T) {
		snp, err := NewSnippet(version, codeSnippet, message, topicID)

		assert.NoError(t, err)
		assert.Equal(t, snp.Version, version)
		assert.Equal(t, snp.CodeSnippet, codeSnippet)
		assert.Equal(t, snp.Message, message)
		assert.Equal(t, snp.TopicID, topicID)
	})

	t.Run("it should be not able to create a snippet with blank code snippet", func(t *testing.T) {
		snp, err := NewSnippet(version, "", message, topicID)

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("codesnippet")).Error(), err.Error())
	})

	t.Run("it should be not able to create a snippet with blank topic id", func(t *testing.T) {
		snp, err := NewSnippet(version, codeSnippet, message, "")

		assert.Error(t, err)
		assert.Nil(t, snp)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("topicid")).Error(), err.Error())
	})
}
