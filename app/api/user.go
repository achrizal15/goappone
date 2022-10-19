package api

import (
	"GoAppOne/app/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
	var db gorm.DB = DbConn()
	param := mux.Vars(r)["id"]
	var users []models.User
	db.Model(&users).Preload("Books").Find(&users, param)
	json.NewEncoder(w).Encode(users)
}
func UserPost(w http.ResponseWriter, r *http.Request) {
	var db gorm.DB = DbConn()
	var name, email string = r.FormValue("name"), r.FormValue("email")
	var user models.User = models.User{Name: name, Email: email}
	err := db.Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var response = map[string]string{}
	response["message"] = "Data submit succesfully"
	json.NewEncoder(w).Encode(response)
	return
}
func UserDelete(w http.ResponseWriter, r *http.Request) {
	var db gorm.DB = DbConn()
	param := mux.Vars(r)["id"]
	err := db.Delete(&models.User{}, param).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var response = map[string]string{"message": "Data deleted succesfully"}
	json.NewEncoder(w).Encode(response)
}
func UserPut(w http.ResponseWriter, r *http.Request) {
	var db gorm.DB = DbConn()
	param := mux.Vars(r)["id"]
	var user *models.User = &models.User{}
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	err := db.Model(&user).Where("id", param).Updates(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.Find(&user, param)

	json.NewEncoder(w).Encode(user)
}
