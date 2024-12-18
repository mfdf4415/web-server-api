package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse request data"})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server create error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user singup successfuly"})

}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "password != email"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "token not generated"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user login successfuly", "token": token})

}
