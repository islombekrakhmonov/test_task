package storage

import (
	"context"
	"task/models"
)



type StorageI interface {
	CloseDB()
	User() UserRepoI
	Phone() PhoneRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUserRequest) (resp *models.CreateUserResponse, err error)
	Login(ctx context.Context, req *models.AuthUserRequest) (resp *models.AuthUserResponse, err error)
	GetByName(ctx context.Context, req *models.GetUserByNameRequest) (resp *models.GetUserByNameResponse, err error)
}

type PhoneRepoI interface {
	Create(ctx context.Context, req *models.CreatePhoneRequest) (resp *models.CreatePhoneResponse, err error)
	CheckDuplicatePhoneNumber(ctx context.Context, phoneNumber string) (bool, error)
	GetByPhone(ctx context.Context, req *models.GetByPhoneRequest) (resp []*models.GetByPhoneResponse, err error)
	Update(ctx context.Context, req *models.UpdatePhoneRequest) (resp *models.UpdatePhoneResponse, err error)
	Delete(ctx context.Context, pKey *models.PhonePKey) (error)
}