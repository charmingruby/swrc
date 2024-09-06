package mongo_repository

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/review/domain/entity"
	"github.com/charmingruby/swrc/internal/review/infra/database/mongo_repository/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCommentVoteMongoRepository(db *mongo.Database) CommentVoteMongoRepository {
	return CommentVoteMongoRepository{
		db: db,
	}
}

type CommentVoteMongoRepository struct {
	db *mongo.Database
}

func (r CommentVoteMongoRepository) FindByID(id string) (entity.CommentVote, error) {
	collection := r.db.Collection(COMMENT_VOTE_COLLECTION)

	filter := bson.D{{Key: "_id", Value: id}}

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return entity.CommentVote{}, core.NewNotFoundErr("comment vote")
	}

	var mongoVote mapper.MongoCommentVote
	if err := res.Decode(&mongoVote); err != nil {
		return entity.CommentVote{}, err
	}

	comment := mapper.MongoCommentVoteToDomain(mongoVote)

	return comment, nil
}

func (r CommentVoteMongoRepository) FindByCommentIDAndAccountID(commentID, accountID string) (entity.CommentVote, error) {
	collection := r.db.Collection(COMMENT_VOTE_COLLECTION)

	filter := bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "account_id", Value: accountID}},
			bson.D{{Key: "comment_id", Value: commentID}},
		}},
	}

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return entity.CommentVote{}, core.NewNotFoundErr("comment vote")
	}

	var mongoVote mapper.MongoCommentVote
	if err := res.Decode(&mongoVote); err != nil {
		return entity.CommentVote{}, err
	}

	vote := mapper.MongoCommentVoteToDomain(mongoVote)

	return vote, nil
}

func (r CommentVoteMongoRepository) Store(vote entity.CommentVote) error {
	collection := r.db.Collection(COMMENT_VOTE_COLLECTION)

	mongoVote := mapper.DomainCommentVoteToMongo(vote)

	if _, err := collection.InsertOne(context.Background(), mongoVote); err != nil {
		return err
	}

	return nil
}

func (r CommentVoteMongoRepository) Delete(vote entity.CommentVote) error {
	collection := r.db.Collection(COMMENT_VOTE_COLLECTION)

	filter := bson.D{{Key: "_id", Value: vote.ID}}

	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}
