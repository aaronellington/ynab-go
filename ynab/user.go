package ynab

import (
	"context"
	"net/http"
)

type UserService struct {
	client *client
}

func (s UserService) Get(ctx context.Context) (UserResponse, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"/user",
		http.NoBody,
	)
	if err != nil {
		return UserResponse{}, err
	}

	target := UserResponse{}

	if err := s.client.Do(request, &target); err != nil {
		return UserResponse{}, err
	}

	return target, nil
}
