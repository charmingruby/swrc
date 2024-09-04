package mapper

import (
	"time"

	"github.com/charmingruby/swrc/internal/review/domain/entity"
)

func DomainSnippetTopicToMongo(topic entity.SnippetTopic) MongoSnippetTopic {
	return MongoSnippetTopic{
		ID:                topic.ID,
		Title:             topic.Title,
		Description:       topic.Description,
		Status:            topic.Status,
		BestAnswerID:      topic.BestAnswerID,
		SnippetSolutionID: topic.SnippetSolutionID,
		CurrentVersion:    topic.CurrentVersion,
		AccountID:         topic.AccountID,
		CreatedAt:         topic.CreatedAt,
	}
}

func MongoSnippetTopicToDomain(topic MongoSnippetTopic) entity.SnippetTopic {
	return entity.SnippetTopic{
		ID:                topic.ID,
		Title:             topic.Title,
		Description:       topic.Description,
		Status:            topic.Status,
		BestAnswerID:      topic.BestAnswerID,
		SnippetSolutionID: topic.SnippetSolutionID,
		CurrentVersion:    topic.CurrentVersion,
		AccountID:         topic.AccountID,
		CreatedAt:         topic.CreatedAt,
	}
}

type MongoSnippetTopic struct {
	ID                string    `json:"id" bson:"_id"`
	Title             string    `json:"title" bson:"title"`
	Description       string    `json:"description" bson:"description"`
	Status            string    `json:"status" bson:"status"`
	CurrentVersion    int       `json:"current_version" bson:"current_version"`
	BestAnswerID      string    `json:"best_answer_id" bson:"best_answer_id"`
	SnippetSolutionID string    `json:"snippet_solution_id" bson:"snippet_solution_id"`
	AccountID         string    `json:"account_id" bson:"account_id"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
}
