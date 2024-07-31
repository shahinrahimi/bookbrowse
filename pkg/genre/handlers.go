package genre

import (
	"log"
	"net/http"

	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

type Handler struct {
	logger *log.Logger
	store  Storage
}

func NewHandler(logger *log.Logger, store Storage) *Handler {
	return &Handler{logger, store}
}

func (h *Handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	genres, err := h.store.GetGenres()
	if err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "error geting all genres"})
	}
	utils.WriteJSON(rw, http.StatusOK, genres)

}

func (h *Handler) GetSingle(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Create(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {

}
