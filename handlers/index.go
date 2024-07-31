package handlers

import (
	"log"

	"github.com/shahinrahimi/bookbrowse/stores"
)

type Handler struct {
	logger *log.Logger
	store  stores.Storage
}

func NewHandler(logger *log.Logger, store stores.Storage) *Handler {
	return &Handler{logger, store}
}
