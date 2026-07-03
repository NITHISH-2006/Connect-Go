package model

import "time"

type User struct{
	ID string `json:"id"`
	Username string `json:"usrename"`
	Email string `jason:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterRequest struct{
	Username string `json:" username" validation:"required"`
	Email string `json:"email" validation:"required"`
	Password string `json:"password" validation:"required, min=6" "`
}

type LoginRequest struct{
	Email string `json:"json:"email" validation:"required, email"`
	Password string `json:"password" validation:"required"`
}