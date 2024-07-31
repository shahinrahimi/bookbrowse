package middlewares

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shahinrahimi/bookbrowse/models"
	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

func (m *Middleware) ValidateBook(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var b models.Book
		if err := utils.FromJSON(&b, r.Body); err != nil {
			m.logger.Println("error deserializing book", err)
			utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
			return
		}
		// validate book
		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(b); err != nil {
			m.logger.Println("validating book failed", err)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "bad request"})
			return
		}
		// Add the book to context
		ctx := context.WithValue(r.Context(), models.KeyBook{}, b)
		r = r.WithContext(ctx)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
