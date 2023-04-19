package router

import (
	"GORM_Project/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

func init() {
	router.HandleFunc("/users", controllers.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/create/user", controllers.Createuser).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", controllers.FindUser).Methods(http.MethodGet)
}

func Run(port string) {
	fmt.Println(http.ListenAndServe(":"+port, router))
}
