package ynab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type client struct {
	httpClient  *http.Client
	accessToken string
}

func (c *client) Do(request *http.Request, target any) error {
	request.Header.Add(
		"Authorization",
		fmt.Sprintf(
			"Bearer %s",
			c.accessToken,
		),
	)

	request.URL.Scheme = "https"
	request.URL.Host = "api.ynab.com"
	request.URL.Path = "/v1" + request.URL.Path

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode >= http.StatusBadRequest {
		responseErr := &ErrorResponse{
			HTTPResponse: response,
		}

		if err := json.Unmarshal(bodyBytes, responseErr); err != nil {
			return err
		}

		return responseErr
	}

	if target != nil {
		if err := json.Unmarshal(bodyBytes, target); err != nil {
			return err
		}
	}

	return nil
}
