package controllers

import (
	"encoding/json"
	"net/http"


	"raz.zaantu.com/m/v0/domain/dto"
	"raz.zaantu.com/m/v0/usecases"
)

type UserController struct {
	UserInteractor usecases.UserInteractor
}

func NewUserController (userInteractor usecases.UserInteractor) *UserController {
	return &UserController{userInteractor}
}

func (controller *UserController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user dto.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)

	if err != nil {
		payload := make(map[string]string)
		payload["message"] = "invalid payload"
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(payload)
		return
	}
	result,err2 := controller.UserInteractor.CreateUser(user)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err2.Error())
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}


func (controller *UserController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.UserInteractor.FindAll()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err2.Error())
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
