package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aaronellington/ynab-go/ynab"
)

func main() {
	ctx := context.Background()

	y := ynab.New(os.Getenv("YNAB_TOKEN"))

	budgets, err := y.Budgets().List(ctx, true)
	if err != nil {
		log.Fatal(err)
	}

	twelveHoursAgo := time.Now().Add(time.Hour * -12)
	beginningOfMonth := getBeginningOfMonth()

	accountsToReconcile := false
	accountCount := 0

	for _, budgetSummary := range budgets.Data.Budgets {
		response, err := y.Budgets().Get(ctx, budgetSummary.ID, nil)
		if err != nil {
			log.Fatal(err)
		}

		budget := response.Data.Budget

		for _, month := range response.Data.Budget.Months {
			if !month.IsCurrentMonth() {
				continue
			}

			for _, category := range month.Categories {
				if category.Balance < 0 {
					log.Printf("OVER BUDGET: %s - %s", budget.Name, category.Name)
				}
			}
		}

		accountIDLookup := map[ynab.AccountID]ynab.Account{}

		for _, account := range budget.Accounts {
			accountIDLookup[account.ID] = account

			if account.Closed || account.Deleted {
				continue
			}

			accountCount++

			reconciliationTarget := beginningOfMonth
			if account.OnBudget {
				reconciliationTarget = twelveHoursAgo
			}

			if account.LastReconciledAt == nil || account.LastReconciledAt.Before(reconciliationTarget) {
				log.Printf("NEED TO RECONCILE ACCOUNT: %s - %s", budget.Name, account.Name)

				accountsToReconcile = true
			}
		}

		for _, transaction := range budget.Transactions {
			account := accountIDLookup[transaction.AccountID]
			if !account.OnBudget {
				continue
			}

			if transaction.CategoryID != nil {
				continue
			}

			if transaction.TransferAccountID != nil {
				continue
			}

			log.Printf("NEED TO SET CATEGORY ON TRANSACTION IN ACCOUNT: %s - %s", budget.Name, account.Name)
		}
	}

	if !accountsToReconcile {
		log.Printf("All %d accounts have been reconciled recently", accountCount)
	}
}

func getBeginningOfMonth() time.Time {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
}
