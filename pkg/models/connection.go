package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/olaishola05/blog-apis/pkg/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

func ConnectDb() (*MongoInstance, error) {
	cfgs, err := config.LoadConfig()

	if err != nil {
		return nil, fmt.Errorf("config package not found")
	}

	mongoURI := "mongodb://" + cfgs.DB.Username + ":" + cfgs.DB.Password + "@" + cfgs.DB.Host + ":" + cfgs.DB.Port + "/" + cfgs.DB.DbName + "?authSource=admin&directConnection=true"
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	db := client.Database(cfgs.DB.DbName)
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	fmt.Println("Database connected successfully")
	return &mg, nil
}

func init() {
	mongoInstance, err := ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	mgo = *mongoInstance
}

func GetDb() *MongoInstance {
	return &mgo
}
