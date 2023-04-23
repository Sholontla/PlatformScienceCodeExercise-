package redis

import (
	"context"
	"encoding/json"

	"finance_server/internal/domain/entity"
	"fmt"
	"time"
)

func RedisChache() {
	var sale []entity.Sale
	var ctxR = context.Background()

	result, errRedis := Cache.Get(ctxR, "daily_revenue").Result()

	if errRedis != nil {

		fmt.Println(errRedis.Error())

		bytes, err := json.Marshal(sale)
		if err != nil {
			panic(err)
		}
		if err := Cache.Set(ctxR, "daily_revenue", bytes, 30*time.Minute).Err(); err != nil {
			panic(err)
		}

	} else {
		json.Unmarshal([]byte(result), &sale)
	}
}

func RedisGetChache() string {
	var ctxR = context.Background()

	result, _ := Cache.Get(ctxR, "daily_revenue").Result()
	return result
}
