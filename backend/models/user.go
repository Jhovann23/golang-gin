package models

import "time"

type User struct {
	Id uint			`json:"id" gorm:"primaryKey"`
	Name string  	`json:"name"`
	Username string	`json:"username" gorm:"unique; not null"`
	Email string	`json:"email" gorm:"unique; not null"`
	Password string	`json:"password"`
	Created_at time.Time	`json:"created_at"`
	Updated_at time.Time	`json:"updated_at"`
}
