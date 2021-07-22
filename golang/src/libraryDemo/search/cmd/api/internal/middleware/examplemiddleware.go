package middleware

import (
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("example middle")
		next(w, r)
	}
}
