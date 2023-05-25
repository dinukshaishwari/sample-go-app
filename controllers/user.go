package controllers

import (
	"log"
	"net/http"

	"sample-go-app/forms"
	"sample-go-app/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	log.Println("Retrieve user")
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User found!", "user": user})
		return
	} else {
		users, err := userModel.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Users found!", "users": users})
		return
	}
}

func (u UserController) Create(ctx *gin.Context) {
	var req forms.UserSignup
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := userModel.Signup(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created!", "users": account})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
