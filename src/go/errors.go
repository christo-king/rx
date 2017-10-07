package main

import (
	"errors"
	"log"
	"net/http"
)

var httpError = http.Error

type HttpError interface {
	error
	Code() int
	LogError() error
}

type HttpResponse struct {
	code     int
	err      error
	logError error
}

func (h HttpResponse) Code() int {
	return h.code
}

func (h HttpResponse) Error() string {
	return h.err.Error()
}

func (h HttpResponse) LogError() error {
	return h.logError
}

type ErrorHandledFunc func(w http.ResponseWriter, r *http.Request) HttpError

type HttpErrorHandler struct {
	wrapped ErrorHandledFunc
}

func (h *HttpErrorHandler) HandleHttpErrors(w http.ResponseWriter, r *http.Request) {
	hr := h.wrapped(w, r)
	if hr.Code() > 299 || hr.Code() < 200 {
		httpError(w, hr.Error(), hr.Code())
	}
	logError := hr.LogError()
	if logError != nil {
		log.Fatal("Error ", hr.Code(), ": ", hr.Error(), ": ", hr.LogError())
	}
}

func NewHttpError(code int, err string) HttpResponse {
	return HttpResponse{code: code, err: errors.New(err)}
}

func NewLogHttpError(code int, err string, logError error) HttpResponse {
	return HttpResponse{code, errors.New(err), logError}
}

func HttpOK() HttpResponse {
	return HttpResponse{200, nil, nil}
}
