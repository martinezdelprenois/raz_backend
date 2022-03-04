package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

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
	userInteractor := usecases.NewUserInteractor(userRepo)
	userController := controllers.NewUserController(userInteractor)
	return *userController
}

func getEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})
	var err error
	dbHandler, err = db.NewDBHandler(getEnvVariable("DB_CONNECTION"), getEnvVariable("DB"))
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}
	userController := getUserController()
	httpRouter.POST("/users", userController.Add)
	httpRouter.GET("/users", userController.FindAll)
	httpRouter.SERVE(getEnvVariable("PORT"))
}
