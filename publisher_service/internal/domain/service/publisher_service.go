package service

import (
	"context"
	"fmt"
	"publisher_service/config"
	"publisher_service/internal/domain/entity"
	clientgrpc "publisher_service/internal/infrastructure/client_grpc"
	"publisher_service/internal/infrastructure/datasource/mongodb"
	loggs "publisher_service/pkg/logger"
	"time"
)

type StoreService struct {
	g  clientgrpc.GrpcClient
	c  config.Config
	db mongodb.MongoService
}

func (s StoreService) CreateOrderService(order entity.Order) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	logger, buffer := loggs.CreateLogger()
	logger.Info("InsertOne func running ...")

	db, coll, _ := s.c.MongoDBConfig()

	conn := s.db.ConnMongoDB().Database(db)
	col := conn.Collection(coll)

	_, errInsert := col.InsertOne(ctx, order)
	if errInsert != nil {
		logger.Error("InsertOne func running ...")
		loggE := loggs.Log{Encoded: buffer.String()}
		fmt.Println(loggE.Encoded, errInsert)
	}
}

func (s StoreService) StreamFinanceOrderService(order entity.Order) entity.Order {
	s.g.GrcpClient(order)
	return order
}
