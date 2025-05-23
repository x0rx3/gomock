package transport

import (
	"gomock/internal/model"
	"net/http"
)

type ResponsePreparer interface {
	Prepare(req *http.Request) (*model.SetResponse, error)
}

type RequestMatcher interface {
	Match(req *http.Request) bool
}

type Router interface {
	Path() string
	Method() string
	Handler() http.Handler
}
