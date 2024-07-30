package genre

import (
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
	store  Storage
}

func NewHandler(logger *log.Logger, store Storage) *Handler {
	return &Handler{logger, store}
}

func (h *Handler) GetAll(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetSingle(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Create(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {

}
