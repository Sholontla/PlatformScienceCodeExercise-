package mongodb

import (
	"context"
	"log"
	"publisher_service/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	conf config.Config
}

func (c MongoService) ConnMongoDB() *mongo.Client {
	_, _, url := c.conf.MongoDBConfig()
	cOptions := options.Client().ApplyURI(url).SetMaxPoolSize(200)

	client, err := mongo.Connect(context.TODO(), cOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("MongoDB Connected ...")
	return client
}

// func (o DAObjectService) CheckConection() int {
// 	err := o..Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 		return 0
// 	}
// 	return 1
// }
