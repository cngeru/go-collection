package controllers

import (
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	UsersController = usersController{}
)

type usersController struct{}

func (controller usersController) GetAll(c *gin.Context) {
	users, getErr := services.UserService.GetAll()
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (controller usersController) Get(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpErr := models.BadRequestError("Invalid ID")
		c.JSON(httpErr.Code, httpErr)
		return
	}
	user, getErr := services.UserService.Get(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller usersController) Create(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		httpErr := models.BadRequestError("Invalid Json")
		c.JSON(httpErr.Code, httpErr)
		return
	}

	createdUser, err := services.UserService.Create(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}
