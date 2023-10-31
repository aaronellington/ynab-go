package ynab

import (
	"context"
	"fmt"
	"net/http"
)

type BudgetsService struct {
	client *client
}

// Returns budgets list with summary information. https://api.ynab.com/v1#/Budgets/getBudgets
func (s *BudgetsService) List(
	ctx context.Context,
	includeAccounts bool,
) (BudgetSummaryResponse, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/budgets?include_accounts=%t", includeAccounts),
		http.NoBody,
	)
	if err != nil {
		return BudgetSummaryResponse{}, err
	}

	target := BudgetSummaryResponse{}

	if err := s.client.Do(request, &target); err != nil {
		return BudgetSummaryResponse{}, err
	}

	return target, nil
}
