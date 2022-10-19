package api

import (
	"GoAppOne/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func BookGet(w http.ResponseWriter, r *http.Request) {
	var rdc redis.Client = RedisConn()
	var book []models.Book
	param := mux.Vars(r)["id"]
	keyRedis := fmt.Sprintf("book%s", param)
	val, err := rdc.Get(keyRedis).Result()
	if err != nil {
		var db gorm.DB = DbConn()
		db.Find(&book, param)
		res, err := json.Marshal(book)
		err = rdc.Set(keyRedis, res, time.Minute*1).Err()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("From MYSQL")
		json.NewEncoder(w).Encode(book)
		return
	}

	fmt.Println("From Redis")
	fmt.Fprint(w, val)
	return
}
func BookPost(w http.ResponseWriter, r *http.Request) {
	var db gorm.DB = DbConn()
	var rdc redis.Client = RedisConn()

	var name string = r.FormValue("name")
	authorId, err := strconv.ParseUint(r.FormValue("authorId"), 10, 64)
	var book models.Book = models.Book{Name: name, AuthorID: uint(authorId)}
	err = db.Create(&book).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// DELETE CACHE REDIS
	rdc.Del("book")
	var response = map[string]string{}
	response["message"] = "Data submit succesfully"
	json.NewEncoder(w).Encode(response)
	return
}
func BookDelete(w http.ResponseWriter, r *http.Request) {
	var db gorm.DB = DbConn()
	var rdc redis.Client = RedisConn()
	param := mux.Vars(r)["id"]
	err := db.Delete(&models.Book{}, param).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	keyRedis := fmt.Sprintf("book%s", param)
	rdc.Del(keyRedis)
	var response = map[string]string{"message": "Data deleted succesfully"}
	json.NewEncoder(w).Encode(response)
}
func BookPut(w http.ResponseWriter, r *http.Request) {
	var db gorm.DB = DbConn()
	var rdc redis.Client = RedisConn()
	param := mux.Vars(r)["id"]
	var book *models.Book = &models.Book{}
	authorId, err := strconv.ParseUint(r.FormValue("authorId"), 10, 64)
	book.Name = r.FormValue("name")
	book.AuthorID = uint(authorId)
	err = db.Model(&book).Where("id", param).Updates(&book).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	keyRedis := fmt.Sprintf("book%s", param)
	rdc.Del(keyRedis)
	db.Find(&book, param)
	json.NewEncoder(w).Encode(book)
}
