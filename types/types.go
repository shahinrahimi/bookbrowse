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
	Data       models.Authors `json:"data"`
	Page       int            `json:"page"`
	TotalPages int            `json:"total_pages"`
}

type PaginatedGenresResponse struct {
	Data       models.Genres `json:"data"`
	Page       int           `json:"page"`
	TotalPages int           `json:"total_pages"`
}

type PaginatedBooksResponse struct {
	Data       models.Books `json:"data"`
	Page       int          `json:"page"`
	TotolPages int          `json:"total_pages"`
}
