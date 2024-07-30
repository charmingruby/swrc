package entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/common/core"
)

const (
	SNIPPET_REACTION_LIKE        = "LIKE"
	SNIPPET_REACTION_CHALLENGING = "CHALLENGING"
	SNIPPET_REACTION_MINDBLOW    = "MIND BLOW"
	SNIPPET_REACTION_BUG         = "BUG"
)

func NewSnippetReaction(reaction, snippetID, accountID string) (*SnippetReaction, error) {
	r := SnippetReaction{
		ID:        core.NewID(),
		Reaction:  reaction,
		SnippetID: snippetID,
		AccountID: accountID,
		CreatedAt: time.Now(),
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return &r, nil
}

type SnippetReaction struct {
	ID        string    `json:"id" validate:"required"`
	Reaction  string    `json:"reaction" validate:"required"`
	SnippetID string    `json:"snippet_id" validate:"required"`
	AccountID string    `json:"account_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}

func (r *SnippetReaction) Validate() error {
	if err := core.ValidateStruct(r); err != nil {
		return err
	}

	if err := r.validateReaction(r.Reaction); err != nil {
		return err
	}

	return nil
}

func (r *SnippetReaction) validateReaction(key string) error {
	reactions := map[string]string{
		SNIPPET_REACTION_LIKE:        SNIPPET_REACTION_LIKE,
		SNIPPET_REACTION_CHALLENGING: SNIPPET_REACTION_CHALLENGING,
		SNIPPET_REACTION_MINDBLOW:    SNIPPET_REACTION_MINDBLOW,
		SNIPPET_REACTION_BUG:         SNIPPET_REACTION_BUG,
	}

	_, ok := reactions[key]
	if !ok {
		return core.NewValidationErr("invalid reaction")
	}

	return nil
}
