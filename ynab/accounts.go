package ynab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountsService struct {
	client *client
}

// Returns all accounts. https://api.ynab.com/v1#/Accounts/getAccounts
func (s *AccountsService) List(
	ctx context.Context,
	budgetID BudgetID,
	lastKnowledgeOfServer *int64,
) (AccountsResponse, error) {
	path := fmt.Sprintf("/budgets/%s/accounts", budgetID)
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
		return AccountsResponse{}, err
	}

	target := AccountsResponse{}

	if err := s.client.Do(request, &target); err != nil {
		return AccountsResponse{}, err
	}

	return target, nil
}

// Creates a new account. https://api.ynab.com/v1#/Accounts/createAccount
func (s *AccountsService) Create(
	ctx context.Context,
	budgetID BudgetID,
	account SaveAccount,
) (AccountResponse, error) {
	payloadBytes, err := json.Marshal(PostAccountWrapper{
		Account: account,
	})
	if err != nil {
		return AccountResponse{}, err
	}

	path := fmt.Sprintf("/budgets/%s/accounts", budgetID)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		path,
		bytes.NewBuffer(payloadBytes),
	)
	if err != nil {
		return AccountResponse{}, err
	}

	target := AccountResponse{}

	if err := s.client.Do(request, &target); err != nil {
		return AccountResponse{}, err
	}

	return target, nil
}

// Returns a single account. https://api.ynab.com/v1#/Accounts/getAccountById
func (s *AccountsService) Get(
	ctx context.Context,
	budgetID BudgetID,
	accountID AccountID,
) (AccountResponse, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/budgets/%s/accounts/%s", budgetID, accountID),
		http.NoBody,
	)
	if err != nil {
		return AccountResponse{}, err
	}

	target := AccountResponse{}

	if err := s.client.Do(request, &target); err != nil {
		return AccountResponse{}, err
	}

	return target, nil
}
