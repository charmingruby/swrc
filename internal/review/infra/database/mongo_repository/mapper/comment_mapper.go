package mapper

import (
	"time"

	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func DomainCommentToMongo(comment entity.Comment) MongoComment {
	return MongoComment{
		ID:              comment.ID,
		Content:         comment.Content,
		ParentCommentID: comment.ParentCommentID,
		SnippetTopicID:  comment.SnippetTopicID,
		AccountID:       comment.AccountID,
		CreatedAt:       comment.CreatedAt,
	}
}

func MongoCommentToDomain(comment MongoComment) entity.Comment {
	return entity.Comment{
		ID:              comment.ID,
		Content:         comment.Content,
		ParentCommentID: comment.ParentCommentID,
		SnippetTopicID:  comment.SnippetTopicID,
		AccountID:       comment.AccountID,
		CreatedAt:       comment.CreatedAt,
	}
}

type MongoComment struct {
	ID              string    `json:"id" bson:"_id"`
	AccountID       string    `json:"account_id" bson:"account_id"`
	Content         string    `json:"content" bson:"content"`
	ParentCommentID string    `json:"parent_comment_id" bson:"parent_comment_id"`
	SnippetTopicID  string    `json:"snippet_topic_id" bson:"snippet_topic_id"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
}
