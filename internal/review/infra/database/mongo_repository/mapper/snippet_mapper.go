package mapper

import (
	"time"

	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func DomainSnippetToMongo(snippet entity.Snippet) MongoSnippet {
	return MongoSnippet{
		ID:          snippet.ID,
		Version:     snippet.Version,
		CodeSnippet: snippet.CodeSnippet,
		Message:     snippet.Message,
		TopicID:     snippet.TopicID,
		CreatedAt:   snippet.CreatedAt,
	}
}

func MongoSnippetToDomain(snippet MongoSnippet) entity.Snippet {
	return entity.Snippet{
		ID:          snippet.ID,
		Version:     snippet.Version,
		CodeSnippet: snippet.CodeSnippet,
		Message:     snippet.Message,
		TopicID:     snippet.TopicID,
		CreatedAt:   snippet.CreatedAt,
	}
}

type MongoSnippet struct {
	ID          string    `json:"id" bson:"_id"`
	Version     int       `json:"version" bson:"version"`
	CodeSnippet string    `json:"code_snippet" bson:"code_snippet"`
	Message     string    `json:"message" bson:"message"`
	TopicID     string    `json:"topic_id" bson:"topic_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
