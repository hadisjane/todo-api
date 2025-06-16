package controller

import (
	"TodoApp/models"
	"TodoApp/service"
	"TodoApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var u models.UserRegister

	if err := c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	if err := service.CreateUser(u); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})

}

func Login(c *gin.Context) {
	var u models.UserLogin

	if err := c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	user, err := service.GetUserByUsernameAndPassword(u.Username, u.Password)
	if err != nil {
		HandleError(c, err)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
