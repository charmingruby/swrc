package review_entity

import (
	"testing"

	"github.com/charmingruby/swrc/internal/core"
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

	t.Run("it should be not able to create a developer metrics without account id", func(t *testing.T) {
		dm, err := NewDeveloperMetrics("")

		assert.Error(t, err)
		assert.Nil(t, dm)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("accountid")).Error(), err.Error())
	})
}
