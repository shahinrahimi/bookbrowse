package handlers

import (
	"net/http"
)

func (h *Handler) GetAllBooks(rw http.ResponseWriter, r *http.Request) {

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
