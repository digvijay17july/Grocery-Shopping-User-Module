package handler

import (

	"Grocery-Shopping-User-Module/src/app/model"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	var result = authenticate(w, r)
// 	if result!=true{
// 		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
// 		return
// 	}
 users:=[] model.User{}
 db.Set("gorm:auto_preload", true).Find(&users)
 respondJSON(w,http.StatusOK,users)
}
func CreateUser(db *gorm.DB,w http.ResponseWriter, r *http.Request){
// 	var result = authenticate(w, r)
// 	if result!=true{
// 		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
// 		return
// 	}
	user:= model.User{}
	decoder := json.NewDecoder(r.Body)
	if err:=decoder.Decode(&user); err!=nil{
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	user.Password=createHash(user.Password)
	user.Username=user.Email
	err :=db.Create(&user)
	if  err.Error!=nil{
		fmt.Println(err.Error.Error())
		respondError(w, http.StatusBadRequest, err.Error.Error())
		return
	}
	jsonUser:= model.NewJSONUser(user)
	respondJSON(w, http.StatusCreated,jsonUser)
}
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	var result = authenticate(w, r)
// 	if result!=true{
// 		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
// 		return
// 	}
	vars := mux.Vars(r)

	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err!=nil{
		fmt.Println(err.Error())
		respondError(w, http.StatusBadRequest, err.Error())
	}
	user := getUserOr404(db, uint(i), w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}
func getUserOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.User {
// 	authenticate(w,r)
	user := model.User{}
	if err := db.Set("gorm:auto_preload", true).First(&user,id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}