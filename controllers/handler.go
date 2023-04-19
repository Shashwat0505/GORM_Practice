package controllers

import (
	"GORM_Project/models"
	"GORM_Project/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	services.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func Createuser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	services.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var user models.User
	services.DB.Find(&user, id)
	json.NewEncoder(w).Encode(user)

}
