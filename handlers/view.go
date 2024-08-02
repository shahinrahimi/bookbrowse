package handlers

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/shahinrahimi/bookbrowse/views/home"
)

type ViewHandler struct {
	logger *log.Logger
}

func NewViewHandler(logger *log.Logger) *ViewHandler {
	return &ViewHandler{logger}
}

func render(rw http.ResponseWriter, r *http.Request, c templ.Component) error {

	return c.Render(r.Context(), rw)
}

func (vh *ViewHandler) HandleHome(rw http.ResponseWriter, r *http.Request) {
	if err := render(rw, r, home.Index()); err != nil {
		vh.logger.Printf("error handle home view: %v", err)
	}
}
