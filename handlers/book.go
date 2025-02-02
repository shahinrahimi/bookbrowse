package handlers

import (
	"database/sql"
	"net/http"

	"github.com/shahinrahimi/bookbrowse/models"
	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

func (h *Handler) GetAllBooks(rw http.ResponseWriter, r *http.Request) {

	// parse query parameters
	page := utils.ParseQueryParamsPage(r)
	limit := utils.ParseQueryParamsLimit(r)
	offset := (page - 1) * limit

	// parse query parameted
	genre_id := utils.ParseQueryParamsGenreID(r)
	author_id := utils.ParseQueryParamsGenreID(r)
	var bs *models.Books
	var totalPages int
	if genre_id > 0 {
		// count all books with specific genre
		count, err := h.store.GetBooksCountWithGenreID(genre_id)
		if err != nil {
			utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
			return
		}
		// calculate totalpages
		totalPages = (count + limit - 1) / limit

		// check page
		if page > totalPages {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: types.NOTFOUND_ERROR})
			return
		}
		bs, err = h.store.GetLimitedBooksByGenreID(genre_id, limit, offset)
		if err != nil {
			utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
			return
		}

	} else if author_id > 0 {
		// count all books
		count, err := h.store.GetBooksCountWithAuthorID(author_id)
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

		bs, err = h.store.GetLimitedBooksByAuthorID(author_id, limit, offset)
		if err != nil {
			utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
			return
		}

	} else {

		// count all books
		count, err := h.store.GetBooksCount()
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

		bs, err = h.store.GetLimitedBooks(limit, offset)
		if err != nil {
			utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: types.INTERNAL_ERROR})
			return
		}
	}

	response := types.PaginatedBooksResponse{
		Data:       *bs,
		Page:       page,
		TotolPages: totalPages,
	}
	utils.WriteJSON(rw, http.StatusOK, response)

}

func (h *Handler) GetSingleBook(rw http.ResponseWriter, r *http.Request) {
	id := utils.ParseIDVars(r)
	if id < 1 {
		utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "invalid id"})
		return
	}

	b, err := h.store.GetBook(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(rw, http.StatusNotFound, types.ApiError{Error: "not found"})
			return
		}
		utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
		return
	}

	utils.WriteJSON(rw, http.StatusOK, b)

}

func (h *Handler) PostBook(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) PutBook(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteBook(rw http.ResponseWriter, r *http.Request) {

}
