package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joao-gabriel-cruz/debora-api/config"
	"github.com/joao-gabriel-cruz/debora-api/controller"
	"github.com/joao-gabriel-cruz/debora-api/model"
	prisma_repository "github.com/joao-gabriel-cruz/debora-api/repositories/prisma"
	usecases "github.com/joao-gabriel-cruz/debora-api/use-cases"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db, _ := config.ConnectDatabase()
	userRepository := prisma_repository.NewUserPrismaRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/users", userController.GetUser)

	server.PATCH("/users/:id", func(c *gin.Context) {
		var user model.User
		c.BindJSON(&user)
		userController.UpdateUser(c, user)
	})
	server.DELETE("/users/:id", userController.DeleteUser)

	server.POST("/users", func(c *gin.Context) {
		var user model.User
		c.BindJSON(&user)
		userController.CreateUser(c, user)
	})

	server.GET("/users/:id", userController.GetUserByID)

	server.Run(":8080")
}
