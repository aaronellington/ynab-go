package ynab

import "net/http"

type ErrorResponse struct {
	ErrorDetail  ErrorDetail    `json:"error"`
	HTTPResponse *http.Response `json:"_"`
}

type ErrorDetail struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func (err ErrorResponse) Error() string {
	return err.ErrorDetail.Name // TODO: improve this
}
