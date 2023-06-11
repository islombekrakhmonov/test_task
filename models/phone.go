package models

import (
	"time"
)


type Phone struct {
	PhoneId string `json:"phone_id"`
	UserId string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`
	IsFax  bool `json:"is_fax"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} 


type PhonePKey struct {
	PhoneId string `json:"phone_id"`
}

type CreatePhoneRequest struct {
	UserId string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`
	IsFax  bool `json:"is_fax"`
}

type CreatePhoneResponse struct {
	PhoneId string `json:"phone_id"`
	UserId string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`
	IsFax  bool `json:"is_fax"`
}


type GetByPhoneRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type GetByPhoneResponse struct {
	UserId string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`
	IsFax  bool `json:"is_fax"`
}

type UpdatePhoneRequest struct {
	PhoneId string `json:"phone_id"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`
	IsFax  bool `json:"is_fax"`
}

type UpdatePhoneResponse struct {
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`
	IsFax  bool `json:"is_fax"`
}