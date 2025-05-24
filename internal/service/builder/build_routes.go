package builder

import (
	"gomock/internal/model"
	"gomock/internal/service/matcher"
	"gomock/internal/service/preparer"
	"gomock/internal/transport"
	"gomock/internal/transport/handler"
	"gomock/internal/transport/route"

	"go.uber.org/zap"
)

// Build is a function type that takes a logger and a template as input
// and returns a slice of transport.Router.
type Build func(log *zap.Logger, template *model.HandlerTamplate) []transport.Router

// BuildRoutes is a function that implements the Build function type.
// It takes a logger and a template as input and returns a slice of transport.Router.
// It is responsible for building routes for the HTTP server based on the provided template.
// The function iterates over the cases in the template and creates request matchers
// and response providers for each case.
var BuildRoutes Build = func(log *zap.Logger, template *model.HandlerTamplate) []transport.Router {
	var r = make([]transport.Router, 0)

	for method, cas := range template.Cases {
		reqResM := make(map[transport.RequestMatcher]transport.ResponsePreparer)
		for _, c := range cas {
			reqResM[matcher.NewRequestMatcher(log, &c.MatchRequestTemplate)] = preparer.NewResponsePreparer(c.SetResponseTemplate)
		}

		r = append(
			r,
			route.New(template.Path, method, handler.New(log, reqResM)),
		)
	}

	return r
}
