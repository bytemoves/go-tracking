package storages

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	*redis.Client
}

const key = "drivers"

var once sync.Once 
var redisClient *RedisClient
func GetRedisClient() *RedisClient{
	once.Do(func(){
		client := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password:  "", // no password set
			DB: 0,       
		})

		redisClient = &RedisClient{client}
	})

	_,err := redisClient.Ping().Result()
	if err  != nil {
		log.Fatalf("could not connect to redis %v",err)
	}

	return redisClient

}


func (c *RedisClient) AddDriverLocation(lng,lat float64, id string){
	
	c.GeoAdd(
		key,
		&redis.GeoLocation{Longitude:lng,Latitude:lat,Name: id  },
	)
}



func  (c *RedisClient) RemoveDriverLocation (id string){
	c.ZRem(key , id)
}
 

func  (c *RedisClient) SearchDrivers (limit int,lat,lng, r float64) []redis.GeoLocation{
	

	res,_ := c.GeoRadius(key,lng,lat,&redis.GeoRadiusQuery{
		Radius: r,
		Unit:   "km",
		WithGeoHash: true,
		WithCoord:true,
		WithDist: true,
		Count: limit,
		Sort: "ASC",
	}).Result()

	return res
}


