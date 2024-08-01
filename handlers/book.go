package handlers

import (
	"net/http"

	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

func (h *Handler) GetAllBooks(rw http.ResponseWriter, r *http.Request) {

	// parse query parameters
	page := utils.ParseQueryParamsPage(r)
	limit := utils.ParseQueryParamsLimit(r)
	offset := (page - 1) * limit

	bs, err := h.store.GetLimitedBooks(limit, offset)
	if err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}

	count, err := h.store.GetBooksCount()
	if err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}

	totalPages := (count + limit - 1) / limit

	if page > totalPages {
		utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: "Not Found"})
		return
	}

	response := types.PaginatedBooksResponse{
		Data:       *bs,
		Page:       page,
		TotolPages: totalPages,
	}
	utils.WriteJSON(rw, http.StatusOK, response)
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
