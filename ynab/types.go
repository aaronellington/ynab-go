package ynab

import "time"

type (
	BudgetID    string
	AccountID   string
	AccountType string
	PayeeID     string
)

const (
	AccountTypeChecking       AccountType = "checking"
	AccountTypeSavings        AccountType = "savings"
	AccountTypeCash           AccountType = "cash"
	AccountTypeCreditCard     AccountType = "creditCard"
	AccountTypeLineOfCredit   AccountType = "lineOfCredit"
	AccountTypeOtherAsset     AccountType = "otherAsset"
	AccountTypeOtherLiability AccountType = "otherLiability"
	AccountTypeMortgage       AccountType = "mortgage"
	AccountTypeAutoLoan       AccountType = "autoLoan"
	AccountTypeStudentLoan    AccountType = "studentLoan"
	AccountTypePersonalLoan   AccountType = "personalLoan"
	AccountTypeMedicalDebt    AccountType = "medicalDebt"
	AccountTypeOtherDebt      AccountType = "otherDebt"
)

type BudgetSummaryResponse struct {
	Data BudgetSummaryResponseData `json:"data"`
}

type BudgetSummaryResponseData struct {
	Budgets       []BudgetSummary `json:"budgets"`
	DefaultBudget BudgetSummary   `json:"default_budget"`
}

type BudgetSummary struct {
	ID             BudgetID        `json:"id"`
	Name           string          `json:"name"`
	LastModifiedOn time.Time       `json:"last_modified_on"`
	FirstMonth     string          `json:"first_month"`
	LastMonth      string          `json:"last_month"`
	DateFormat     *DateFormat     `json:"date_format"`
	CurrencyFormat *CurrencyFormat `json:"currency_format"`
	Accounts       []Account       `json:"accounts"`
}

type Account struct {
	ID                  BudgetID                  `json:"id"`
	Name                string                    `json:"name"`
	Type                AccountType               `json:"type"`
	OnBudget            bool                      `json:"on_budget"`
	Closed              bool                      `json:"closed"`
	Note                *string                   `json:"note"`
	Balance             int64                     `json:"balance"`
	ClearedBalance      int64                     `json:"cleared_balance"`
	UnclearedBalance    int64                     `json:"uncleared_balance"`
	TransferPayeeID     *PayeeID                  `json:"transfer_payee_id"`
	DirectImportLinked  bool                      `json:"direct_import_linked"`
	DirectImportInError bool                      `json:"direct_import_in_error"`
	LastReconciledAt    *time.Time                `json:"last_reconciled_at"`
	DebtOriginalBalance *int64                    `json:"debt_original_balance"`
	DebtInterestRates   *LoadAccountPeriodicValue `json:"debt_interest_rates"`
	DebtMinimumPayments *LoadAccountPeriodicValue `json:"debt_minimum_payments"`
	DebtEscrowAmounts   *LoadAccountPeriodicValue `json:"debt_escrow_amounts"`
	Deleted             bool                      `json:"deleted"`
}

type DateFormat struct {
	Format string `json:"format"`
}

type CurrencyFormat struct {
	Description      string `json:"description"`
	IsoCode          string `json:"iso_code"`
	ExampleFormat    string `json:"example_format"`
	DecimalDigits    int32  `json:"decimal_digits"`
	DecimalSeparator string `json:"decimal_separator"`
	SymbolFirst      bool   `json:"symbol_first"`
	GroupSeparator   string `json:"group_separator"`
	CurrencySymbol   string `json:"currency_symbol"`
	DisplaySymbol    bool   `json:"display_symbol"`
}

type LoadAccountPeriodicValue struct{}
