package parser

import "net/http"

type Parser interface {
	Parse(r *http.Request, data interface{}) error
	GetType() string
}
