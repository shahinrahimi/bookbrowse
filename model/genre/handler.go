package genre

import "log"

type Handler struct {
	logger *log.Logger
	store  Storage
}

func NewHandler(logger *log.Logger, store Storage) *Handler {
	return &Handler{logger, store}
}
