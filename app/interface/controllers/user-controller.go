package controllers

import (
	"encoding/json"
	"net/http"
	"../../domain/dto"
	"../../usercases"
)

type UserController struct {
	UserInteractor usecases.UserInteractor
}

func NewUserController (userInteractor usecases.UserInteractor) *UserController {
	return &UserController{userInteractor}
}

func (controller *UserController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().set("Content-Type", "application/json")
	var user dto.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.userInteractor.CreateUser(user)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *UserController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.userInteractor.FindAll()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}