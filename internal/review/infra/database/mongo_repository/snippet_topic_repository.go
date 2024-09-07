package mongo_repository

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/internal/review/infra/database/mongo_repository/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSnippetTopicMongoRepository(db *mongo.Database) SnippetTopicMongoRepository {
	return SnippetTopicMongoRepository{
		db: db,
	}
}

type SnippetTopicMongoRepository struct {
	db *mongo.Database
}

func (r SnippetTopicMongoRepository) FindByID(id string) (entity.SnippetTopic, error) {
	collection := r.db.Collection(SNIPPET_TOPIC_COLLECTION)

	filter := bson.D{{Key: "_id", Value: id}}

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return entity.SnippetTopic{}, core.NewNotFoundErr("snippet topic")
	}

	var mongoTopic mapper.MongoSnippetTopic
	if err := res.Decode(&mongoTopic); err != nil {
		return entity.SnippetTopic{}, err
	}

	topic := mapper.MongoSnippetTopicToDomain(mongoTopic)

	return topic, nil
}

func (r SnippetTopicMongoRepository) Store(topic entity.SnippetTopic) error {
	collection := r.db.Collection(SNIPPET_TOPIC_COLLECTION)

	mongoTopic := mapper.DomainSnippetTopicToMongo(topic)

	if _, err := collection.InsertOne(context.Background(), mongoTopic); err != nil {
		return err
	}

	return nil
}

func (r SnippetTopicMongoRepository) Delete(id string) error {
	collection := r.db.Collection(SNIPPET_TOPIC_COLLECTION)

	filter := bson.D{{Key: "_id", Value: id}}

	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}

func (r SnippetTopicMongoRepository) FindMany(id, status, accountID string) ([]entity.SnippetTopic, error) {
	collection := r.db.Collection(SNIPPET_TOPIC_COLLECTION)

	filter := bson.D{}

	if id != "" || status != "" || accountID != "" {
		filter = bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "_id", Value: id}},
				bson.D{{Key: "status", Value: status}},
				bson.D{{Key: "account_id", Value: accountID}},
			}},
		}
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var mongoTopics []mapper.MongoSnippetTopic
	if err = cursor.All(context.Background(), &mongoTopics); err != nil {
		return nil, err
	}

	topics := []entity.SnippetTopic{}
	for _, t := range mongoTopics {
		topics = append(topics, mapper.MongoSnippetTopicToDomain(t))
	}

	return topics, nil
}

func (r SnippetTopicMongoRepository) Save(topic entity.SnippetTopic) error {
	collection := r.db.Collection(SNIPPET_TOPIC_COLLECTION)

	filter := bson.D{{Key: "_id", Value: topic.ID}}

	update := bson.M{"$set": bson.M{
		"title":               topic.Title,
		"description":         topic.Description,
		"status":              topic.Status,
		"current_version":     topic.CurrentVersion,
		"best_answer_id":      topic.BestAnswerID,
		"snippet_solution_id": topic.SnippetSolutionID,
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
