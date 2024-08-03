package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joao-gabriel-cruz/debora-api/model"
	usecases "github.com/joao-gabriel-cruz/debora-api/use-cases"
)

type UserController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(usecase usecases.UserUseCase) UserController {
	return UserController{
		userUseCase: usecase,
	}
}

func (u *UserController) CreateUser(c *gin.Context, user model.User) {
	err := u.userUseCase.CreateUser(c, user)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		c.JSON(200, gin.H{
			"message": "User created",
		})
	}
}

func (u *UserController) UpdateUser(c *gin.Context, user model.User) {
	id := c.Param("id")
	err := u.userUseCase.UpdateUser(c, id, user)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		c.JSON(200, gin.H{
			"message ": "User updated ",
		})
	}
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := u.userUseCase.DeleteUser(c, id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "User deleted",
	})
}

func (u *UserController) GetUser(c *gin.Context) {
	result, err := u.userUseCase.GetUser(c)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"users": result,
	})
}

func (u *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	result, err := u.userUseCase.GetUserByID(c, id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"user": result,
	})
}
