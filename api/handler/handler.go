package handler

import (
	"task/config"
	"task/storage"
)


type Handler struct {
	cfg      *config.Config
	storages storage.StorageI
}


func NewHandler(cfg *config.Config, store storage.StorageI) *Handler {
	return &Handler{
		cfg:      cfg,
		storages: store,
	}
}