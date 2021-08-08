package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisDB *redis.Client

func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func redisExample() {
	err := redisDB.Set("score", 98, time.Second * 100).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	result, err := redisDB.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}

	fmt.Println("score: ", result)

	val2, err := redisDB.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

func redisExample2()  {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{
			Score:  98,
			Member: "Golang",
		},
		redis.Z{
			Score:  99,
			Member: "Java",
		},
		redis.Z{
			Score:  96,
			Member: "Python",
		},
		redis.Z{
			Score:  95,
			Member: "C++",
		},
	}

	// ZAdd
	num, err := redisDB.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := redisDB.ZIncrBy(zsetKey, 10, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	results, err := redisDB.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, v := range results {
		fmt.Println(v.Member, v.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	rets, err := redisDB.ZRangeByScore(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, v := range rets {
		fmt.Println(v)
	}

}

func main()  {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
		return
	}
	fmt.Printf("connect redis success\n")

	redisExample2()
	
}
