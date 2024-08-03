package mongo_repository

import (
	"context"

	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/account/infra/database/mongo_repository/mapper"
	"github.com/charmingruby/swrc/internal/common/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAccountMongoRepository(db *mongo.Database) AccountMongoRepository {
	return AccountMongoRepository{
		db: db,
	}
}

type AccountMongoRepository struct {
	db *mongo.Database
}

func (r *AccountMongoRepository) FindByID(id string) (entity.Account, error) {
	collection := r.db.Collection(accountCollection)

	filter := bson.D{{Key: "_id", Value: id}}

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return entity.Account{}, core.NewNotFoundErr("account")
	}

	var mongoAcc mapper.MongoAccount
	if err := res.Decode(&mongoAcc); err != nil {
		return entity.Account{}, err
	}

	acc := mapper.MongoAccountToDomain(mongoAcc)

	return acc, nil
}

func (r *AccountMongoRepository) FindByEmail(email string) (entity.Account, error) {
	collection := r.db.Collection(accountCollection)

	filter := bson.D{{Key: "email", Value: email}}

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return entity.Account{}, core.NewNotFoundErr("account")
	}

	var mongoAcc mapper.MongoAccount
	if err := res.Decode(&mongoAcc); err != nil {
		return entity.Account{}, err
	}

	acc := mapper.MongoAccountToDomain(mongoAcc)

	return acc, nil
}

func (r *AccountMongoRepository) FindByGithubDisplayName(githubDisplayName string) (entity.Account, error) {
	collection := r.db.Collection(accountCollection)

	filter := bson.D{{Key: "github_display_name", Value: githubDisplayName}}

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return entity.Account{}, core.NewNotFoundErr("account")
	}

	var mongoAcc mapper.MongoAccount
	if err := res.Decode(&mongoAcc); err != nil {
		return entity.Account{}, err
	}

	acc := mapper.MongoAccountToDomain(mongoAcc)

	return acc, nil
}

func (r *AccountMongoRepository) Store(acc *entity.Account) error {
	return nil
}

func (r *AccountMongoRepository) SaveVerification(acc *entity.Account) error {
	return nil
}

func (r *AccountMongoRepository) SaveRole(acc *entity.Account) error {
	return nil
}