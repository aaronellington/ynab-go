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

	minimumReconciledTimeForOnBudget := time.Now().Add(time.Hour * -24 * 2)
	minimumReconciledTimeForOffBudget := time.Now().Add(time.Hour * -24 * 7)

	for _, budget := range budgets.Data.Budgets {
		for _, account := range budget.Accounts {
			if account.OnBudget {
				if account.LastReconciledAt.Before(minimumReconciledTimeForOnBudget) {
					log.Printf("Need to reconcile account: %s - %s", budget.Name, account.Name)
				}
			} else {
				if account.LastReconciledAt.Before(minimumReconciledTimeForOffBudget) {
					log.Printf("Need to reconcile account: %s - %s", budget.Name, account.Name)
				}
			}
		}
	}
}
