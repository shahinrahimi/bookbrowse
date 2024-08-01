package handlers

import (
	"net/http"
	"strconv"

	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

func (h *Handler) GetAllBooks(rw http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 100
	// parse query parameters
	if p := r.URL.Query().Get("page"); p != "" {
		parsePage, err := strconv.Atoi(p)
		if err == nil && parsePage > 0 {
			page = parsePage
		}
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	offset := (page - 1) * limit

	bs, err := h.store.GetLimitedBooks(limit, offset)
	if err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, bs)
}

func (h *Handler) GetSingleBook(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) PostBook(rw http.ResponseWriter, r *http.Request) {
	// var book models.Book
	// if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
	// 	utils.WriteJSON(rw, http.StatusInternalServerError, err)
	// 	return
	// }
}

func (h *Handler) PutBook(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteBook(rw http.ResponseWriter, r *http.Request) {

}
