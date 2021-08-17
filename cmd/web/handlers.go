package main

import (
	"net/http"

	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const uri = "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func InitialMigration() {

}

func getUsers(w http.ResponseWriter, r *http.Request) {

}
func getUser(w http.ResponseWriter, r *http.Request) {

}
func createUser(w http.ResponseWriter, r *http.Request) {

}
func updateUser(w http.ResponseWriter, r *http.Request) {

}
func deleteUser(w http.ResponseWriter, r *http.Request) {

}
