package main

import (
	"fmt"
	"net/http"
	"time"

	"go-test/service"

	"github.com/gorilla/mux"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	fmt.Println("http server running on localhost:8080")

	// with mux
	r := mux.NewRouter()
	r.HandleFunc("/register", userSvc.RegisterHandler)
	r.HandleFunc("/user", userSvc.GetUserHandler)
	r.HandleFunc("/user/{id}", userSvc.GetUserHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()
}
