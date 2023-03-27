package application

import (
	"context"
	"net/http"
)

var _ Process = (*GoRestful)(nil)

type GoRestful struct{}

func (s *GoRestful) Start(_ context.Context) {
	http.ListenAndServe(":8080", nil)
}

func (s *GoRestful) Stop(_ context.Context) {

}
