package ynab

import (
	"net/http"
	"time"
)

func New(
	accessToken string,
	configFuncs ...ConfigFunc,
) *Service {
	c := &client{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
		accessToken: accessToken,
	}

	s := &Service{
		client: c,
		user: &UserService{
			client: c,
		},
		budgets: &BudgetsService{
			client: c,
		},
		accounts: &AccountsService{
			client: c,
		},
		categories: &CategoriesService{
			client: c,
		},
		scheduledTransactions: &ScheduledTransactionsService{
			client: c,
		},
		months: &MonthsService{
			client: c,
		},
		payees: &PayeesService{
			client: c,
		},
		payeeLocations: &PayeeLocationsService{
			client: c,
		},
		transactions: &TransactionsService{
			client: c,
		},
	}

	for _, configFunc := range configFuncs {
		configFunc(s)
	}

	return s
}

type Service struct {
	client                *client
	user                  *UserService
	budgets               *BudgetsService
	accounts              *AccountsService
	categories            *CategoriesService
	payees                *PayeesService
	payeeLocations        *PayeeLocationsService
	months                *MonthsService
	transactions          *TransactionsService
	scheduledTransactions *ScheduledTransactionsService
}

func (s *Service) User() *UserService {
	return s.user
}

func (s *Service) Budgets() *BudgetsService {
	return s.budgets
}

func (s *Service) Accounts() *AccountsService {
	return s.accounts
}

func (s *Service) Categories() *CategoriesService {
	return s.categories
}

func (s *Service) Payees() *PayeesService {
	return s.payees
}

func (s *Service) PayeeLocations() *PayeeLocationsService {
	return s.payeeLocations
}

func (s *Service) Months() *MonthsService {
	return s.months
}

func (s *Service) Transactions() *TransactionsService {
	return s.transactions
}

func (s *Service) ScheduledTransactions() *ScheduledTransactionsService {
	return s.scheduledTransactions
}
