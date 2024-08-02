package handlers

import (
	"database/sql"
	"net/http"

	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

func (h *Handler) GetAllAuthors(rw http.ResponseWriter, r *http.Request) {
	// parse query parameters
	page := utils.ParseQueryParamsPage(r)
	limit := utils.ParseQueryParamsLimit(r)
	offset := (page - 1) * limit

	// count all authors
	count, err := h.store.GetAuthorsCount()
	if err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}

	// calculate totalpages
	totalPages := (count + limit - 1) / limit

	// check page
	if page > totalPages {
		utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
		return
	}

	as, err := h.store.GetLimitedAuthors(limit, offset)
	if err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}

	response := types.PaginatedAuthorsResponse{
		Data:       *as,
		Page:       page,
		TotalPages: totalPages,
	}
	utils.WriteJSON(rw, http.StatusOK, response)
}

func (h *Handler) GetSingleAuthor(rw http.ResponseWriter, r *http.Request) {
	id := utils.ParseIDVars(r)
	if id < 1 {
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: types.INVALID_ID_ERROR})
		return
	}

	a, err := h.store.GetAuthor(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}

	utils.WriteJSON(rw, http.StatusOK, a)

}

func (h *Handler) PostAuthor(rw http.ResponseWriter, r *http.Request) {
	// var book Book
	// if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
	// 	utils.WriteJSON(rw, http.StatusInternalServerError, err)
	// 	return
	// }
}

func (h *Handler) PutAuthor(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteAuthor(rw http.ResponseWriter, r *http.Request) {

}
