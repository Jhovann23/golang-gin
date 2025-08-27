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

func Register(c *gin.Context) {
	//inisialisasi menangkap request user
	var req = structs.UserCreateRequest{}

	//validasi user menggunakan binding dari gin
	if err := c.ShouldBindJSON(&req); err != nil {
		//jika validasi gagal kembalikan respon error json menggunakan structs error 
		 c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validasi Error",
			Errors: helper.TranslateErrorMessage(err),
		 })
		 return
	}

	//buat data user dari request struct
	user := models.User{
		Name: req.Name,
		Username: req.Username,
		Email: req.Email,
		Password: helper.HashPasswords(req.Password),
	}

	//simpan data user ke database
	if err := database.DB.Create(&user).Error; err != nil{
		//cek apakah ada data yang sama menggunakan. contoh username/email sudah terdaftar
		if helper.IsDuplicateEntryError(err) {
			//jika duplikat, kasih response error dengan http status 409 error conflict
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate Entry Error",
				Errors: helper.TranslateErrorMessage(err),
			})			
		} else{
			//jika error lain, kirimkan http status error 500 internal server error
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed To Create User",
				Errors: helper.TranslateErrorMessage(err),
			})
		}
		return
	}
	//jika berhasil, kirimkan responses sukses
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User Created Successfully",
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

