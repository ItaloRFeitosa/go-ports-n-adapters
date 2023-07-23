package httpmodel

import "net/http"

type StatusOK struct{}

func (StatusOK) StatusCode() int {
	return http.StatusOK
}

type StatusCreated struct{}

func (StatusCreated) StatusCode() int {
	return http.StatusCreated
}

type StatusNoContent struct{}

func (StatusNoContent) StatusCode() int {
	return http.StatusNoContent
}

type StatusNotFound struct{}

func (StatusNotFound) StatusCode() int {
	return http.StatusNotFound
}
