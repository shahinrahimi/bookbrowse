package handlers

import "net/http"

func (h *Handler) GetAllAuthors(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetSingleAuthor(rw http.ResponseWriter, r *http.Request) {

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
