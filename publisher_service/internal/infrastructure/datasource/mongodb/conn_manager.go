package mongodb

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"store_service/config"
// 	"store_service/internal/domain/entity"
// 	loggs "store_service/pkg/logger"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type DAObjectService struct {
// 	conifg config.Config
// }

// func (o DAObjectService) InsertOne(client *mongo.Client, request entity.Order) *mongo.InsertOneResult {
// 	logger, buffer := loggs.CreateLogger()
// 	logger.Info("InsertOne func running ...")
// 	logg := loggs.Log{Encoded: buffer.String()}
// 	fmt.Println(logg.Encoded)
// 	db, coll, _ := o.conifg.MongoDBConfig()
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()
// 	conn := client.Database(db)
// 	col := conn.Collection(coll)
// 	response, errInsert := col.InsertOne(ctx, request)

// 	if errInsert != nil {
// 		logger.Error("InsertOne func running ...")
// 		loggE := loggs.Log{Encoded: buffer.String()}
// 		fmt.Println("**************", loggE.Encoded, errInsert)
// 	}
// 	return response
// }

// func (o DAObjectService) GetAll(client *mongo.Client) ([]map[string]interface{}, error) {
// 	//db, collection, _ := o.conifg.MongoDBConfig()

// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()
// 	var responseSlice []map[string]interface{}

// 	conn := client.Database("stores")
// 	coll := conn.Collection("stores")

// 	cursor, err := coll.Find(ctx, responseSlice)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	for cursor.Next(ctx) {
// 		result := make(map[string]interface{})
// 		err := cursor.Decode(&result)
// 		if err != nil {
// 			return nil, err
// 		}
// 		responseSlice = append(responseSlice, result)
// 	}
// 	defer cancel()

// 	return responseSlice, nil
// }
