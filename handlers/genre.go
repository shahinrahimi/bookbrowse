package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/shahinrahimi/bookbrowse/models"
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
	g := r.Context().Value(models.KeyGenre{}).(models.Genre)
	g.Name = strings.ToLower(g.Name)
	if err := h.store.CreateGenre(&g); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			utils.WriteJSON(rw, http.StatusConflict, types.ApiError{Error: types.DUPLICATION_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}
	utils.WriteJSON(rw, http.StatusCreated, types.ApiSuccess{Message: "New genre created"})
}

func (h *Handler) PutGenre(rw http.ResponseWriter, r *http.Request) {
	id := utils.ParseIDVars(r)
	if id < 1 {
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: types.INVALID_ID_ERROR})
		return
	}
	// check genre id is exists
	if _, err := h.store.GetGenre(id); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}
	g := r.Context().Value(models.KeyGenre{}).(models.Genre)
	g.Name = strings.ToLower(g.Name)
	if err := h.store.UpdateGenre(id, &g); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			utils.WriteJSON(rw, http.StatusConflict, types.ApiError{Error: types.DUPLICATION_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.NOTFOUND_ERROR})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: fmt.Sprintf("Genre with %d updated", id)})
}

func (h *Handler) DeleteGenre(rw http.ResponseWriter, r *http.Request) {
	id := utils.ParseIDVars(r)
	if id < 1 {
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: types.INVALID_ID_ERROR})
		return
	}
	// check genre id is exists
	if _, err := h.store.GetGenre(id); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}

	// TODO
	// fetch books with genre id
	// delete all records in genres_books for combination of book id and genre id

	if err := h.store.DeleteGenre(id); err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: fmt.Sprintf("Author with %d deleted", id)})
}
