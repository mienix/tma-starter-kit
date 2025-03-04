package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/devflex-pro/tma-starter-kit/backend/config"
	"github.com/devflex-pro/tma-starter-kit/backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func New(
	ctx context.Context,
	conf config.MongoUserDBConfig,
) (*MongoUserDB, error) {
	clientOpts := options.Client().
		ApplyURI(conf.URI)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(conf.DBName)
	collection := db.Collection(conf.CollectionName)

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.
		Indexes().
		CreateOne(ctx, indexModel)
	if err != nil {
		return nil, err
	}

	return &MongoUserDB{
		client:     client,
		collection: collection,
	}, nil
}

func (m *MongoUserDB) Save(
	ctx context.Context,
	user domain.User,
) error {
	filter := bson.M{"id": user.ID}
	update := bson.M{
		"$set": user,
		"$setOnInsert": bson.M{
			"created_at": time.Now(),
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err := m.collection.
		UpdateOne(ctx, filter, update, opts)
	return err
}

func (m *MongoUserDB) Read(
	ctx context.Context,
	id int,
) (domain.User, error) {
	filter := bson.M{"id": id}
	var user domain.User
	err := m.collection.
		FindOne(ctx, filter).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, err
}

func (m *MongoUserDB) Close(
	ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
