package handler

import (
	"Grocery-Shopping/src/app/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
 users:=[] model.User{}
 db.Find(&users)
 respondJSON(w,http.StatusOK,users)
}
func CreateUser(db *gorm.DB,w http.ResponseWriter, r *http.Request){
	user:= model.User{}
	decoder := json.NewDecoder(r.Body)
	if err:=decoder.Decode(&user); err!=nil{
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	err :=db.Create(&user)
	if  err.Error!=nil{
		fmt.Println(err.Error.Error())
		respondError(w, http.StatusBadRequest, err.Error.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	respondJSON(w, http.StatusOK, employee)
}
// getUser404 gets a employee instance if exists, or respond the 404 error otherwise
func getEmployeeOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.User {
	employee := model.User{}
	if err := db.First(&employee, model.User{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &employee
}