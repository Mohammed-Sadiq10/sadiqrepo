package main

import (
	"log"
	"sadiq/Go_Rest_API/controller"
	"sadiq/Go_Rest_API/database"
	"sadiq/Go_Rest_API/repository"
	"sadiq/Go_Rest_API/service"
	"sadiq/Go_Rest_API/validation"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db := database.DbConection()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	// Middleware for header validation
	router.Use(validation.ValidateHeaders())

	// Routes
	router.POST("/users/v1/add-user", userController.CreateUser)
	router.GET("/users/v1/get-users", userController.GetAllUsers)
	router.GET("/users/v1/get-user/:id", userController.GetUserByID)
	router.GET("/users/v1/user-email/:email", userController.GetUserByEmail)
	router.PUT("/users/v1/update-user/:id", userController.UpdateUser)
	router.DELETE("/users/v1/remove-user/:id", userController.DeleteUser)

	// Run the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server")
	}
	defer db.Close()
}
