package middlewares

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shahinrahimi/bookbrowse/models"
	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

func (m *Middleware) ValidateGenre(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var g models.Genre
		if err := utils.FromJSON(&g, r.Body); err != nil {
			m.logger.Println("error deserializing genre", err)
			utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
			return
		}
		// validate book
		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(g); err != nil {
			m.logger.Println("validating genre failed", err)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "bad request"})
			return
		}
		// Add the book to context
		ctx := context.WithValue(r.Context(), models.KeyGenre{}, g)
		r = r.WithContext(ctx)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)

	})
}
