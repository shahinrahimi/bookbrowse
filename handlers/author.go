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
	a := r.Context().Value(models.KeyAuthor{}).(models.Author)
	a.Name = strings.ToLower(a.Name)
	if err := h.store.CreateAuthor(&a); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			utils.WriteJSON(rw, http.StatusConflict, types.ApiError{Error: types.DUPLICATION_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}
	utils.WriteJSON(rw, http.StatusCreated, types.ApiSuccess{Message: "New author created"})
}
func (h *Handler) PutAuthor(rw http.ResponseWriter, r *http.Request) {
	id := utils.ParseIDVars(r)
	if id < 1 {
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: types.INVALID_ID_ERROR})
		return
	}
	// check author id is exists
	if _, err := h.store.GetAuthor(id); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}
	a := r.Context().Value(models.KeyAuthor{}).(models.Author)
	a.Name = strings.ToLower(a.Name)
	if err := h.store.UpdateAuthor(id, &a); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			utils.WriteJSON(rw, http.StatusConflict, types.ApiError{Error: types.DUPLICATION_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.NOTFOUND_ERROR})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: fmt.Sprintf("Author with %d updated", id)})
}
func (h *Handler) DeleteAuthor(rw http.ResponseWriter, r *http.Request) {
	id := utils.ParseIDVars(r)
	if id < 1 {
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: types.INVALID_ID_ERROR})
		return
	}
	// check author id is exists
	if _, err := h.store.GetAuthor(id); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}

	// TODO
	// delete all books for author

	// TODO
	// fetch genres id for book and
	// delete all records in genres_books with combinatuions of book id and genres

	if err := h.store.DeleteAuthor(id); err != nil {
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
		return
	}
	utils.WriteJSON(rw, http.StatusOK, types.ApiSuccess{Message: fmt.Sprintf("Author with %d deleted", id)})
}
