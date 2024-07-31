package middlewares

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shahinrahimi/bookbrowse/models"
	"github.com/shahinrahimi/bookbrowse/types"
	"github.com/shahinrahimi/bookbrowse/utils"
)

func (m *Middleware) ValidateAuthor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var a models.Author
		if err := utils.FromJSON(&a, r.Body); err != nil {
			m.logger.Println("error deserializing author", err)
			utils.WriteJSON(rw, http.StatusInternalServerError, types.ApiError{Error: "internal error"})
			return
		}
		// validate author
		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(a); err != nil {
			m.logger.Println("validating genre failed", err)
			utils.WriteJSON(rw, http.StatusBadRequest, types.ApiError{Error: "bad request"})
			return
		}
		// Add the author to context
		ctx := context.WithValue(r.Context(), models.KeyAuthor{}, a)
		r = r.WithContext(ctx)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)

	})
}
