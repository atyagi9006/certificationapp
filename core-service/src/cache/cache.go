package cache

import (
	"log"
	"strconv"

	"github.com/atyagi9006/certificationapp/core-service/src/config"
	"github.com/go-redis/redis"
)

type Client struct {
	RClient *redis.Client
}

var Config *config.Database

func Init(conf *config.Database) {
	Config = conf
	//client = redisClient(db)
	//defer client.Close()
}

func RedisClient() *Client {
	address := Config.DBConfig.URL + ":" + strconv.Itoa(int(Config.DBConfig.Port))
	dbname, err := strconv.Atoi(Config.DBConfig.DatabaseName)
	if err == nil {
		log.Printf("i=%d, type: %T\n", dbname, dbname)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: Config.DBConfig.Password, // password set
		DB:       dbname,                   // use default DB
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)
	// Output: PONG <nil>

	rclient := Client{
		RClient: client,
	}
	return &rclient
}

// Read func read cache from redis
func (c *Client) Read(key string) string {
	val, err := c.RClient.Get(key).Result()
	if err != nil {
		panic(err)
	}
	if err == redis.Nil {
		log.Println(key, "does not exist")
	}
	log.Println("key", val)
	// Output: key value
	return val
}

//Write func write cache tn redis
func (c *Client) Write(key string, value interface{}) {
	err := c.RClient.Set(key, value, config.RedisTTL).Err()
	if err != nil {
		panic(err)
	}
}
func (c *Client) Delete(key string) {
	n, err := c.RClient.Del(key).Result()
	if err != nil {
		panic(err)
	}
	log.Println("result del ", n)
}
