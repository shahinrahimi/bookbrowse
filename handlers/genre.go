package handlers

import (
	"database/sql"
	"net/http"

	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

func (h *Handler) GetAllGenres(rw http.ResponseWriter, r *http.Request) {
	// parse query parameters
	page := utils.ParseQueryParamsPage(r)
	limit := utils.ParseQueryParamsLimit(r)
	offset := (page - 1) * limit

	// count all genres
	count, err := h.store.GetGenresCount()
	if err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}
	// calculate totalpage
	totalPages := (count + limit - 1) / limit
	if page > totalPages {
		utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
		return
	}

	gs, err := h.store.GetLimitedGenres(limit, offset)
	if err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}

	response := types.PaginatedGenresResponse{
		Data:       *gs,
		Page:       page,
		TotalPages: totalPages,
	}

	utils.WriteJSON(rw, http.StatusOK, response)

}

func (h *Handler) GetSingleGenre(rw http.ResponseWriter, r *http.Request) {

	id := utils.ParseIDVars(r)
	if id < 1 {
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: types.INVALID_ID_ERROR})
		return
	}

	g, err := h.store.GetGenre(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}

	utils.WriteJSON(rw, http.StatusOK, g)

}

func (h *Handler) PostGenre(rw http.ResponseWriter, r *http.Request) {
	// var book Book
	// if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
	// 	utils.WriteJSON(rw, http.StatusInternalServerError, err)
	// 	return
	// }
}

func (h *Handler) PutGenre(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteGenre(rw http.ResponseWriter, r *http.Request) {

}
