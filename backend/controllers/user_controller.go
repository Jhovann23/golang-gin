package controllers

import (
	"backend/database"
	"backend/helper"
	"backend/models"
	"backend/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Data Lists Users",
		Data: users,
	})
}

func CreateUser(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors: helper.TranslateErrorMessage(err),
		})
		return
	}

	//inialisasi user baru
	user := models.User{
		Name: req.Name,
		Username: req.Username,
		Email: req.Email,
		Password: helper.HashPasswords(req.Password),
	}

	//simpen ke database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Errors: helper.TranslateErrorMessage(err),
		})
		return
	}

	//respon success
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Succes create user",
		Data: structs.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}

func FindUserById(c*gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Not found",
			Errors: helper.TranslateErrorMessage(err),
		})
		return
	}

	//sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User found",
		Data: structs.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}
