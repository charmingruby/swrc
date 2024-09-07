package mongo_repository

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/internal/review/infra/database/mongo_repository/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSnippetMongoRepository(db *mongo.Database) SnippetMongoRepository {
	return SnippetMongoRepository{
		db: db,
	}
}

type SnippetMongoRepository struct {
	db *mongo.Database
}

func (r SnippetMongoRepository) Store(snippet entity.Snippet) error {
	collection := r.db.Collection(SNIPPET_COLLECTION)

	mongoSnippet := mapper.DomainSnippetToMongo(snippet)

	if _, err := collection.InsertOne(context.Background(), mongoSnippet); err != nil {
		return err
	}

	return nil
}

func (r SnippetMongoRepository) Save(snippet entity.Snippet) error {
	collection := r.db.Collection(SNIPPET_COLLECTION)

	filter := bson.D{{Key: "_id", Value: snippet.ID}}

	update := bson.M{"$set": bson.M{
		"message":      snippet.Message,
		"code_snippet": snippet.CodeSnippet,
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r SnippetMongoRepository) FindByID(id string) (entity.Snippet, error) {
	collection := r.db.Collection(SNIPPET_COLLECTION)

	filter := bson.D{{Key: "_id", Value: id}}

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return entity.Snippet{}, core.NewNotFoundErr("snippet")
	}

	var mongoSnippet mapper.MongoSnippet
	if err := res.Decode(&mongoSnippet); err != nil {
		return entity.Snippet{}, err
	}

	snippet := mapper.MongoSnippetToDomain(mongoSnippet)

	return snippet, nil
}

func (r SnippetMongoRepository) FindManyByTopicID(topicID string) ([]entity.Snippet, error) {
	collection := r.db.Collection(SNIPPET_COLLECTION)

	filter := bson.D{{}}

	if topicID != "" {
		filter = bson.D{{Key: "snippet_topic_id", Value: topicID}}
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var mongoSnippets []mapper.MongoSnippet
	if err = cursor.All(context.Background(), &mongoSnippets); err != nil {
		return nil, err
	}

	snippets := []entity.Snippet{}
	for _, s := range mongoSnippets {
		snippets = append(snippets, mapper.MongoSnippetToDomain(s))
	}

	return snippets, nil
}

func (r SnippetMongoRepository) DeleteMany(snippets []entity.Snippet) error {
	collection := r.db.Collection(SNIPPET_COLLECTION)

	objectIDs := make([]primitive.ObjectID, len(snippets))
	for i, snippet := range snippets {
		objectID, err := primitive.ObjectIDFromHex(snippet.ID)
		if err != nil {
			return err
		}
		objectIDs[i] = objectID
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	if _, err := collection.DeleteMany(context.Background(), filter); err != nil {
		return err
	}

	return nil
}
