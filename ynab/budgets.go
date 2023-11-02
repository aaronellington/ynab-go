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

// Returns a single budget with all related entities. This resource is effectively a full budget export. https://api.ynab.com/v1#/Budgets/getBudgetById
func (s *BudgetsService) Get(
	ctx context.Context,
	id BudgetID,
	lastKnowledgeOfServer *int64,
) (BudgetDetailResponse, error) {
	path := fmt.Sprintf("/budgets/%s", id)
	if lastKnowledgeOfServer != nil {
		path = fmt.Sprintf(
			"%s?last_knowledge_of_server=%d",
			path,
			lastKnowledgeOfServer,
		)
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		path,
		http.NoBody,
	)
	if err != nil {
		return BudgetDetailResponse{}, err
	}

	target := BudgetDetailResponse{}

	if err := s.client.Do(request, &target); err != nil {
		return BudgetDetailResponse{}, err
	}

	return target, nil
}
