package ynab

import (
	"fmt"
	"net/http"
)

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
	return fmt.Sprintf(
		"YNAB API Error [%s]: %s - %s",
		err.ErrorDetail.ID,
		err.ErrorDetail.Name,
		err.ErrorDetail.Detail,
	)
}
