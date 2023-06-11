package models

import "time"

type User struct {
	UserId string  `json:"user_id"`
	Login   string  `json:"login"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Age      int   `json:"age"`
	Token    string `json:"token"`
	RefreshToken    string `json:"refresh_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Login   string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int  `json:"age"`
}

type CreateUserResponse struct {
	UserId   string `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int   `json:"age"`
}

type AuthUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthUserResponse struct {
	UserId   string `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Token	string `json:"token"`
	RefreshToken	string `json:"refresh_token"`
}


type UserPKey struct {
	UserId string `json:"user_id"`
}

type GetUserByNameRequest struct {
	Name string `json:"name"`
}

type GetUserByNameResponse struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Age      int `json:"age"`
}	