package controllers

import (
	"backend/database"
	"backend/helper"
	"backend/models"
	"backend/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	//struct yang menampung data user
	var req = structs.UserLoginRequest{}
	var user = models.User{}

	//validasi input dari request body menggunakan shouldBinJSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors: helper.TranslateErrorMessage(err),
		}) 
		return
	}


	//cari username di database
	//jika tidak ditemukan error response unauthorized
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors: helper.TranslateErrorMessage(err),
		})
		return
	}

	//bandingkan password user dan password yang ada di database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil{
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Passwords",
			Errors: helper.TranslateErrorMessage(err),
		})
	}

	//jika login berhasil generate token
	token := helper.GenerateToken(user.Username)

	//kirim statusOK jika semua validasi benar
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Successfully",
		Data: structs.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Token: &token,
		},
	})
}
