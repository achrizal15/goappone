package api

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Api struct {
	DBConn *gorm.DB
	APP_NAME,
	APP_PORT,
	DB_USER,
	DB_PASSWORD,
	DB_NAME string
}

func getEnv(env_key string, falback string) (result string) {
	result = os.Getenv(env_key)
	if result == "" {
		return falback
	}
	return result
}
func (api *Api) initializeEnv() {
	godotenv.Load()
	api.APP_NAME = getEnv("APP_NAME", "GoAppOne")
	api.APP_PORT = getEnv("APP_PORT", "localhost:8080")
	api.DB_NAME = getEnv("DB_NAME", "media_visual_banyuwangi")
	api.DB_USER = getEnv("DB_USER", "root")
	api.DB_PASSWORD = getEnv("DB_PASSWORD", "")
}
func (api *Api) databaseConnection() gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", api.DB_USER, api.DB_PASSWORD, api.DB_NAME)
	var err error
	api.DBConn, err = gorm.Open(mysql.New(mysql.Config{
		DisableWithReturning: true,
		DSN:                  dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return *api.DBConn
}
func DbConn() gorm.DB {
	api := Api{}
	api.initializeEnv()
	db := api.databaseConnection()
	return db
}

var ctx = context.Background()

func RedisConn() redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return *rdb

}
