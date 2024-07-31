package middlewares

import "log"

type Middleware struct {
	logger *log.Logger
}

func NewMiddlware(logger *log.Logger) *Middleware {
	return &Middleware{logger}
}
