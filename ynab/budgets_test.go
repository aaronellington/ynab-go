package ynab_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/aaronellington/ynab-go/ynab"
)

type MockRoundTripper struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

func (m MockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(r)
}

func ServeFileAsResponse(
	statusCode int,
	filePath string,
) (*http.Response, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	response := &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(file),
	}

	return response, nil
}

func TestBudgetsListSuccess(t *testing.T) {
	ctx := context.Background()
	accessToken := "fake-access-token"

	y := ynab.New(
		accessToken,
		ynab.SetRoundTripper(MockRoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", accessToken) {
					return nil, errors.New("wrong/missing access token")
				}

				return ServeFileAsResponse(http.StatusOK, "test_files/budgets/list/success.json")
			},
		}),
	)

	budgets, err := y.Budgets().List(ctx, false)
	if err != nil {
		t.Fatal(err)
	}

	if len(budgets.Data.Budgets) != 5 {
		t.Fatalf("Wrong number of budgets found. Found %d, expected %d", len(budgets.Data.Budgets), 5)
	}
}

func TestBudgetsListFailure(t *testing.T) {
	ctx := context.Background()
	accessToken := "fake-access-token"

	y := ynab.New(
		accessToken,
		ynab.SetRoundTripper(MockRoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", accessToken) {
					return nil, errors.New("wrong/missing access token")
				}

				return ServeFileAsResponse(http.StatusUnauthorized, "test_files/unauthorized.json")
			},
		}),
	)

	_, err := y.Budgets().List(ctx, false)
	if err == nil {
		t.Fatal("there should have been an error")
	}

	ynabErr, ok := err.(ynab.ErrorResponse)
	if !ok {
		t.Fatal(err)
	}

	actual := ynabErr.ErrorDetail
	expected := ynab.ErrorDetail{
		ID:     "401",
		Name:   "unauthorized",
		Detail: "Unauthorized",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("wrong error values")
	}
}
