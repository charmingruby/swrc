package review_entity

import (
	"testing"

	"github.com/charmingruby/swrc/internal/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewCommentVote(t *testing.T) {
	commentID := "comment id"
	accountID := "account id"
	isUp := true

	t.Run("it should be able to create a new comment vote", func(t *testing.T) {
		cv, err := NewCommentVote(isUp, accountID, commentID)

		assert.NoError(t, err)
		assert.NotNil(t, cv)
		assert.Equal(t, isUp, cv.IsUp)
		assert.Equal(t, accountID, cv.AccountID)
		assert.Equal(t, commentID, cv.CommentID)
	})

	t.Run("it should be not able to create comment vote without account id", func(t *testing.T) {
		cv, err := NewCommentVote(isUp, "", commentID)

		assert.Error(t, err)
		assert.Nil(t, cv)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("accountid")).Error(), err.Error())
	})

	t.Run("it should be not able to create comment vote without comment id", func(t *testing.T) {
		cv, err := NewCommentVote(isUp, accountID, "")

		assert.Error(t, err)
		assert.Nil(t, cv)
		assert.Equal(t, core.NewValidationErr(core.ErrRequired("commentid")).Error(), err.Error())
	})
}
