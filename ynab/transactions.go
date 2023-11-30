package ynab

import (
	"context"
	"fmt"
	"net/http"
)

type TransactionsService struct {
	client *client
}

func (s TransactionsService) List(ctx context.Context, budgetID BudgetID) (TransactionsResponse, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/budgets/%s/transactions", budgetID),
		http.NoBody,
	)
	if err != nil {
		return TransactionsResponse{}, err
	}

	target := TransactionsResponse{}

	if err := s.client.Do(request, &target); err != nil {
		return TransactionsResponse{}, err
	}

	return target, nil
}
