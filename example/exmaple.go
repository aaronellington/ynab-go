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

	for _, budget := range budgets.Data.Budgets {
		for _, account := range budget.Accounts {
			if account.Closed || account.Deleted {
				continue
			}

			accountCount++

			reconciliationTarget := beginningOfMonth
			if account.OnBudget {
				reconciliationTarget = twelveHoursAgo
			}

			if account.LastReconciledAt == nil || account.LastReconciledAt.Before(reconciliationTarget) {
				log.Printf("Need to reconcile account: %s - %s - https://app.ynab.com/%s/accounts/%s", budget.Name, account.Name, budget.ID, account.ID)

				accountsToReconcile = true
			}
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
