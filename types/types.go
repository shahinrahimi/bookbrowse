package types

import "github.com/shahinrahimi/bookbrowse/models"

type ApiError struct {
	Error string `json:"error"`
}
type ApiSuccess struct {
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type PaginatedAnyResponse struct {
	Data       []any `json:"data"`
	Page       int   `json:"page"`
	TotalPages int   `json:"total_pages"`
}

type PaginatedAuthorsResponse struct {
	Data        []models.Author `json:"data"`
	Page        int             `json:"page"`
	TotoalPages int             `json:"total_pages"`
}

type PaginatedGenresResponse struct {
	Data       []models.Genre `json:"data"`
	Page       int            `json:"page"`
	TotalPages int            `json:"total_pages"`
}

type PaginatedBooksResponse struct {
	Data       models.Books `json:"data"`
	Page       int          `json:"page"`
	TotolPages int          `json:"total_pages"`
}
