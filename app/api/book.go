package api

import (
	"GoAppOne/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func BookGet(w http.ResponseWriter, r *http.Request) {
	var rdc redis.Client = RedisConn()
	var books []models.Book

	val, err := rdc.Get("books").Result()
	if err != nil {
		var db gorm.DB = DbConn()
		db.Find(&books)
		res, err := json.Marshal(books)
		err = rdc.Set("books", res, time.Minute*1).Err()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("From MYSQL")
		json.NewEncoder(w).Encode(res)
		return
	}
	fmt.Println("From Redis")
	fmt.Fprint(w, val)
	return
}
