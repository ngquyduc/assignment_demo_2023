package db

import (
	"context"
	"encoding/json"
	"github.com/ngquyduc/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var db *redis.Client
var ctx = context.Background()
var addr string
var pass string

func init() {
	if os.Getenv("ENV") == "PROD" {
		addr = "redis:6379"
		pass = "redis"
	} else {
		addr = "localhost:6379"
		pass = ""
	}
	db = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err := db.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connected to redis")
	}
}

func LoadMsg(msg *rpc.Message) error {
	jsonObj, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = db.RPush(ctx, msg.Chat, jsonObj).Result()
	if err != nil {
		return err
	}
	return err
}

func GetMsg(key string, cursor, limit int64, reverse bool) ([]*rpc.Message, bool, int64, error) {
	msgJsons, err := db.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, false, 0, err
	}

	msgs := make([]*rpc.Message, 0)

	for _, msgInterface := range msgJsons {
		var msg rpc.Message
		err = json.Unmarshal([]byte(msgInterface), &msg)
		if err != nil {
			return nil, false, 0, err
		}
		msgs = append(msgs, &msg)
	}

	length := len(msgJsons)
	hasMore := true
	if length == 0 {
		hasMore = false
	}

	nextCursor := int64(0)
	if cursor >= 0 && limit >= 0 {
		var filteredMsgs []*rpc.Message
		count := 0
		for i, msg := range msgs {
			if count >= int(limit) {
				break
			}
			if i == length-1 {
				hasMore = false
			}
			if hasMore {
				nextCursor = msgs[i+1].SendTime
			}
			if msg.SendTime >= cursor {
				filteredMsgs = append(filteredMsgs, msg)
				count++
			}
		}
		msgs = filteredMsgs
	}

	// reverse messages
	if reverse {
		for i, j := 0, len(msgs)-1; i < j; i, j = i+1, j-1 {
			msgs[i], msgs[j] = msgs[j], msgs[i]
		}
	}

	if !hasMore {
		nextCursor = 0
	}

	return msgs, hasMore, nextCursor, nil
}
