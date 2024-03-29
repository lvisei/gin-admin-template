package app

import (
	"context"
	"time"

	"gin-admin-template/internal/app/config"
	imongo "gin-admin-template/internal/app/model/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitMongo 初始化mongo存储
func InitMongo() (*mongo.Client, func(), error) {
	cfg := config.C.Mongo
	client, cleanFunc, err := imongo.NewClient(&imongo.Config{
		URI:      cfg.URI,
		Database: cfg.Database,
		Timeout:  time.Duration(cfg.Timeout) * time.Second,
	})
	if err != nil {
		return nil, cleanFunc, err
	}

	err = imongo.CreateIndexes(context.Background(), client)
	if err != nil {
		return nil, cleanFunc, err
	}

	return client, cleanFunc, nil
}
