package mongo

import (
	"context"

	"fixit.com/backend/src/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userRepo struct {
	db         string
	collection string
	client     *mongo.Client
}

func CreateUserRepo(client *mongo.Client, db string, collection string) *userRepo {
	return &userRepo{client: client, db: db, collection: collection}
}

func (r *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	collection := r.client.Database(r.db).Collection(r.collection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	collection := r.client.Database(r.db).Collection(r.collection)
	filter := bson.M{"username": username}
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
