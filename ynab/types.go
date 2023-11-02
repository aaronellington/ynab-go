package ynab

import "time"

type (
	AccountID       string
	AccountType     string
	BudgetID        string
	CategoryGroupID string
	CategoryID      string
	PayeeID         string
	PayeeLocationID string
	GoalType        string
)

const (
	AccountTypeAutoLoan       AccountType = "autoLoan"
	AccountTypeCash           AccountType = "cash"
	AccountTypeChecking       AccountType = "checking"
	AccountTypeCreditCard     AccountType = "creditCard"
	AccountTypeLineOfCredit   AccountType = "lineOfCredit"
	AccountTypeMedicalDebt    AccountType = "medicalDebt"
	AccountTypeMortgage       AccountType = "mortgage"
	AccountTypeOtherAsset     AccountType = "otherAsset"
	AccountTypeOtherDebt      AccountType = "otherDebt"
	AccountTypeOtherLiability AccountType = "otherLiability"
	AccountTypePersonalLoan   AccountType = "personalLoan"
	AccountTypeSavings        AccountType = "savings"
	AccountTypeStudentLoan    AccountType = "studentLoan"
)

const (
	GoalTypeTargetCategoryBalance       GoalType = "TB"
	GoalTypeTargetCategoryBalanceByDate GoalType = "TBD"
	GoalTypeMonthlyFunding              GoalType = "MF"
	GoalTypePlanYourSpending            GoalType = "NEED"
	GoalTypeDebt                        GoalType = "DEBT"
)

type BudgetSummaryResponse struct {
	Data BudgetSummaryResponseData `json:"data"`
}

type BudgetSummaryResponseData struct {
	Budgets       []BudgetSummary `json:"budgets"`
	DefaultBudget BudgetSummary   `json:"default_budget"`
}

type BudgetDetailResponse struct {
	Data BudgetDetailResponseData
}

type BudgetDetailResponseData struct {
	Budget          BudgetDetail `json:"budget"`
	ServerKnowledge int64        `json:"server_knowledge"`
}

type BudgetDetail struct {
	BudgetSummary
	Payees         []Payee         `json:"payees"`
	PayeeLocations []PayeeLocation `json:"payee_locations"`
	CategoryGroups []CategoryGroup `json:"category_groups"`
	Categories     []Category      `json:"categories"`
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

type Payee struct {
	ID                PayeeID    `json:"id"`
	Name              string     `json:"name"`
	TransferAccountID *AccountID `json:"transfer_account_id"`
	Deleted           bool       `json:"deleted"`
}

type PayeeLocation struct {
	ID        PayeeLocationID `json:"id"`
	PayeeID   PayeeID         `json:"payee_id"`
	Latitude  string          `json:"latitude"`
	Longitude string          `json:"longitude"`
	Deleted   bool            `json:"deleted"`
}

type CategoryGroup struct {
	ID      CategoryGroupID `json:"id"`
	Name    string          `json:"name"`
	Hidden  bool            `json:"hidden"`
	Deleted bool            `json:"deleted"`
}

type Category struct {
	ID                     CategoryID      `json:"id"`
	CategoryGroupID        CategoryGroupID `json:"category_group_id"`
	CategoryGroupName      string          `json:"category_group_name"`
	Name                   string          `json:"name"`
	Hidden                 bool            `json:"hidden"`
	Note                   string          `json:"note"`
	Budgeted               int64           `json:"budgeted"`
	Activity               int64           `json:"activity"`
	Balance                int64           `json:"balance"`
	GoalType               *GoalType       `json:"goal_type"`
	GoalDay                *int32          `json:"goal_day"`
	GoalCadence            *int32          `json:"goal_cadence"`
	GoalCadenceFrequency   *int32          `json:"goal_cadence_frequency"`
	GoalCreationMonth      *string         `json:"goal_creation_month"`
	GoalTarget             *int64          `json:"goal_target"`
	GoalTargetMonth        *string         `json:"goal_target_month"`
	GoalPercentageComplete *int32          `json:"goal_percentage_complete"`
	GoalMonthsToBudget     *int32          `json:"goal_months_to_budget"`
	GoalUnderFunded        *int64          `json:"goal_under_funded"`
	GoalOverallFunded      *int64          `json:"goal_overall_funded"`
	GoalOverallLeft        *int64          `json:"goal_overall_left"`
	Deleted                bool            `json:"deleted"`
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
