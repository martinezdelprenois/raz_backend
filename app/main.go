package main

import (
	"fmt"
	"log"
	"net/http"

	"raz.zaantu.com/m/v0/infrastructure/db"

	"raz.zaantu.com/m/v0/usecases"

	"raz.zaantu.com/m/v0/interface/controllers"

	"raz.zaantu.com/m/v0/infrastructure/router"
	"raz.zaantu.com/m/v0/interface/repository"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	dbHandler  db.DBHandler
)

func getUserController() controllers.UserController {
	userRepo := repository.NewUserRepository(dbHandler)
	userInteractor := usecases.NewUserInteractor(userRepository)
	userController := controllers.NewUserController(userInteractor)
	return *userController
}


func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})
	var err error
	dbHandler, err = db.NewDBHandler("mongodb://localhost:27017", "raz")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}
	bookController := getBookController()
	httpRouter.POST("/book", bookController.Add)
	httpRouter.GET("/book", bookController.FindAll)
	httpRouter.SERVE(":8000")
}