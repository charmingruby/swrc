package review_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDeveloperMetrics(t *testing.T) {
	accountID := "account id"

	t.Run("it should be able to create a developer metrics", func(t *testing.T) {
		dm, err := NewDeveloperMetrics(accountID)

		assert.NoError(t, err)
		assert.Equal(t, 0, dm.SnippetsPublished)
		assert.Equal(t, 0, dm.Stars)
		assert.Equal(t, accountID, dm.AccountID)
	})
}
