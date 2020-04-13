package main

import (
	"./app"
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Auth).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")
	port := os.Getenv("port")
	fmt.Println(port)
	err := http.ListenAndServe(":" + port, router)
	fmt.Println(err)
}
