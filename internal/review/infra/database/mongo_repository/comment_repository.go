package mongo_repository

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/internal/review/infra/database/mongo_repository/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCommentMongoRepository(db *mongo.Database) CommentMongoRepository {
	return CommentMongoRepository{
		db: db,
	}
}

type CommentMongoRepository struct {
	db *mongo.Database
}

func (r *CommentMongoRepository) FindByID(id string) (entity.Comment, error) {
	collection := r.db.Collection(COMMENT_COLLECTION)

	filter := bson.D{{Key: "_id", Value: id}}

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return entity.Comment{}, core.NewNotFoundErr("comment")
	}

	var mongoComment mapper.MongoComment
	if err := res.Decode(&mongoComment); err != nil {
		return entity.Comment{}, err
	}

	comment := mapper.MongoCommentToDomain(mongoComment)

	return comment, nil
}

func (r *CommentMongoRepository) FindMany(
	id string,
	accountID string,
	snippetTopicID string,
	parentCommentID string,
) ([]entity.Comment, error) {
	collection := r.db.Collection(COMMENT_COLLECTION)

	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "_id", Value: id}},
			bson.D{{Key: "account_id", Value: accountID}},
			bson.D{{Key: "snippet_topic_id", Value: snippetTopicID}},
			bson.D{{Key: "parent_comment_id", Value: parentCommentID}},
		}},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var mongoComments []mapper.MongoComment
	if err = cursor.All(context.Background(), &mongoComments); err != nil {
		return nil, err
	}

	comments := []entity.Comment{}
	for _, c := range mongoComments {
		comments = append(comments, mapper.MongoCommentToDomain(c))
	}

	return comments, nil
}

func (r *CommentMongoRepository) Store(comment entity.Comment) error {
	collection := r.db.Collection(COMMENT_COLLECTION)

	mongoComment := mapper.DomainCommentToMongo(comment)

	if _, err := collection.InsertOne(context.Background(), mongoComment); err != nil {
		return err
	}

	return nil
}

func (r *CommentMongoRepository) Delete(comment entity.Comment) error {
	collection := r.db.Collection(COMMENT_COLLECTION)

	filter := bson.D{{Key: "_id", Value: comment.ID}}

	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}

func (r *CommentMongoRepository) DeleteManyByParentCommentID(parentCommentID string) error {
	collection := r.db.Collection(COMMENT_COLLECTION)

	filter := bson.D{{Key: "parent_comment_id", Value: parentCommentID}}

	if _, err := collection.DeleteMany(context.Background(), filter); err != nil {
		return err
	}

	return nil
}
