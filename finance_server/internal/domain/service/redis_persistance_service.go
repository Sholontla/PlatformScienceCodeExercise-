package service

import (
	"context"
	"encoding/json"
	"finance_server/internal/domain/entity"
	r "finance_server/internal/infrastructure/persistance/redis"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
}

func (f CacheService) CacheOrders(salesData entity.Sale) (entity.Sale, error) {
	var ctxR = context.Background()

	bytes, err := json.Marshal(salesData)
	if err != nil {
		log.Println("Error while Marshal salesData", err.Error())
	}
	if err := r.Cache.Set(ctxR, "orders", bytes, 60*time.Minute).Err(); err != nil {
		log.Println("Error while Setting data into redis: ", err.Error())
	}
	return salesData, nil
}

func (f CacheService) GetDataToProcess() ([]entity.Sale, error) {
	var data []entity.Sale
	var ctxR = context.Background()

	result, errRedis := r.Cache.Get(ctxR, "orders").Result()
	if errRedis != nil {
		fmt.Println(errRedis.Error())
		return data, errRedis
	} else {
		if err := json.Unmarshal([]byte(result), &data); err != nil {
			fmt.Println("Error while Unmarshalling data: ", err.Error())
			return data, err
		}
	}

	return data, nil
}

func SaveData(data *entity.Sale) error {
	// Get existing data from Redis
	val, err := r.Cache.Get(context.Background(), "orders").Result()
	if err == redis.Nil {
		// If key doesn't exist, initialize data as an empty slice
		val = "[]"
	} else if err != nil {
		// Return error if there was a problem getting the data from Redis
		return fmt.Errorf("failed to get data from Redis: %w", err)
	}

	// Unmarshal existing data into a slice of Data objects
	var dataList []entity.Sale
	if err := json.Unmarshal([]byte(val), &dataList); err != nil {
		// Return error if there was a problem unmarshalling the data
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	// Append new data to the slice
	dataList = append(dataList, *data)

	// Marshal the slice back into JSON
	newData, err := json.Marshal(dataList)
	if err != nil {
		// Return error if there was a problem marshalling the data
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	// Save the new data to Redis
	err = r.Cache.Set(context.Background(), "orders", string(newData), 60*time.Minute).Err()
	if err != nil {
		// Return error if there was a problem saving the data to Redis
		return fmt.Errorf("failed to save data to Redis: %w", err)
	}

	return nil
}
