package ynab

import "time"

type (
	AccountID                 string
	AccountType               string
	Milliunits                int64
	BudgetID                  string
	CategoryGroupID           string
	CategoryID                string
	DebtTransactionType       string
	FlagColor                 string
	GoalType                  string
	ImportID                  string
	PayeeID                   string
	PayeeLocationID           string
	ScheduledSubTransactionID string
	ScheduledTransactionID    string
	ScheduleFrequency         string
	SubTransactionID          string
	TransactionID             string
	TransactionStatus         string
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

const (
	TransactionStatusCleared    TransactionStatus = "cleared"
	TransactionStatusUncleared  TransactionStatus = "uncleared"
	TransactionStatusReconciled TransactionStatus = "reconciled"
)

const (
	FlagColorRed    FlagColor = "red"
	FlagColorOrange FlagColor = "orange"
	FlagColorYellow FlagColor = "yellow"
	FlagColorGreen  FlagColor = "green"
	FlagColorBlue   FlagColor = "blue"
	FlagColorPurple FlagColor = "purple"
)

const (
	DebtTransactionTypePayment           DebtTransactionType = "payment"
	DebtTransactionTypeRefund            DebtTransactionType = "refund"
	DebtTransactionTypeFee               DebtTransactionType = "fee"
	DebtTransactionTypeInterest          DebtTransactionType = "interest"
	DebtTransactionTypeEscrow            DebtTransactionType = "escrow"
	DebtTransactionTypeBalanceAdjustment DebtTransactionType = "balanceAdjustment"
	DebtTransactionTypeCredit            DebtTransactionType = "credit"
	DebtTransactionTypeCharge            DebtTransactionType = "charge"
)

const (
	ScheduleFrequencyNever           ScheduleFrequency = "never"
	ScheduleFrequencyDaily           ScheduleFrequency = "daily"
	ScheduleFrequencyWeekly          ScheduleFrequency = "weekly"
	ScheduleFrequencyEveryOtherWeek  ScheduleFrequency = "everyOtherWeek"
	ScheduleFrequencyTwiceAMonth     ScheduleFrequency = "twiceAMonth"
	ScheduleFrequencyEvery4Weeks     ScheduleFrequency = "every4Weeks"
	ScheduleFrequencyMonthly         ScheduleFrequency = "monthly"
	ScheduleFrequencyEveryOtherMonth ScheduleFrequency = "everyOtherMonth"
	ScheduleFrequencyEvery3Months    ScheduleFrequency = "every3Months"
	ScheduleFrequencyEvery4Months    ScheduleFrequency = "every4Months"
	ScheduleFrequencyTwiceAYear      ScheduleFrequency = "twiceAYear"
	ScheduleFrequencyYearly          ScheduleFrequency = "yearly"
	ScheduleFrequencyEveryOtherYear  ScheduleFrequency = "everyOtherYear"
)

type Account struct {
	ID                  BudgetID                  `json:"id"`
	Name                string                    `json:"name"`
	Type                AccountType               `json:"type"`
	OnBudget            bool                      `json:"on_budget"`
	Closed              bool                      `json:"closed"`
	Note                *string                   `json:"note"`
	Balance             Milliunits                `json:"balance"`
	ClearedBalance      Milliunits                `json:"cleared_balance"`
	UnclearedBalance    Milliunits                `json:"uncleared_balance"`
	TransferPayeeID     *PayeeID                  `json:"transfer_payee_id"`
	DirectImportLinked  bool                      `json:"direct_import_linked"`
	DirectImportInError bool                      `json:"direct_import_in_error"`
	LastReconciledAt    *time.Time                `json:"last_reconciled_at"`
	DebtOriginalBalance *Milliunits               `json:"debt_original_balance"`
	DebtInterestRates   *LoadAccountPeriodicValue `json:"debt_interest_rates"`
	DebtMinimumPayments *LoadAccountPeriodicValue `json:"debt_minimum_payments"`
	DebtEscrowAmounts   *LoadAccountPeriodicValue `json:"debt_escrow_amounts"`
	Deleted             bool                      `json:"deleted"`
}

type AccountResponse struct {
	Data AccountResponseData `json:"data"`
}

type AccountResponseData struct {
	Account Account `json:"account"`
}

type AccountsResponse struct {
	Data AccountsResponseData `json:"data"`
}

type AccountsResponseData struct {
	Accounts        []Account `json:"accounts"`
	ServerKnowledge int64     `json:"server_knowledge"`
}

type BudgetDetail struct {
	BudgetSummary
	Payees                   []Payee                       `json:"payees"`
	PayeeLocations           []PayeeLocation               `json:"payee_locations"`
	CategoryGroups           []CategoryGroup               `json:"category_groups"`
	Categories               []Category                    `json:"categories"`
	Months                   []MonthDetail                 `json:"months"`
	Transactions             []TransactionSummary          `json:"transactions"`
	SubTransactions          []SubTransaction              `json:"subtransactions"`
	ScheduledTransactions    []ScheduledTransactionSummary `json:"scheduled_transactions"`
	ScheduledSubTransactions []ScheduledSubTransaction     `json:"scheduled_subtransactions"`
}

type BudgetDetailResponse struct {
	Data BudgetDetailResponseData
}

type BudgetDetailResponseData struct {
	Budget          BudgetDetail `json:"budget"`
	ServerKnowledge int64        `json:"server_knowledge"`
}

type BudgetSettings struct {
	DateFormat     *DateFormat     `json:"date_format"`
	CurrencyFormat *CurrencyFormat `json:"currency_format"`
}

type BudgetSettingsResponse struct {
	Data BudgetSettingsResponseData `json:"data"`
}

type BudgetSettingsResponseData struct {
	Settings BudgetSettings `json:"settings"`
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

type BudgetSummaryResponse struct {
	Data BudgetSummaryResponseData `json:"data"`
}

type BudgetSummaryResponseData struct {
	Budgets       []BudgetSummary `json:"budgets"`
	DefaultBudget BudgetSummary   `json:"default_budget"`
}

type Category struct {
	ID                     CategoryID      `json:"id"`
	CategoryGroupID        CategoryGroupID `json:"category_group_id"`
	CategoryGroupName      string          `json:"category_group_name"`
	Name                   string          `json:"name"`
	Hidden                 bool            `json:"hidden"`
	Note                   string          `json:"note"`
	Budgeted               Milliunits      `json:"budgeted"`
	Activity               Milliunits      `json:"activity"`
	Balance                Milliunits      `json:"balance"`
	GoalType               *GoalType       `json:"goal_type"`
	GoalDay                *int32          `json:"goal_day"`
	GoalCadence            *int32          `json:"goal_cadence"`
	GoalCadenceFrequency   *int32          `json:"goal_cadence_frequency"`
	GoalCreationMonth      *string         `json:"goal_creation_month"`
	GoalTarget             *Milliunits     `json:"goal_target"`
	GoalTargetMonth        *string         `json:"goal_target_month"`
	GoalPercentageComplete *int32          `json:"goal_percentage_complete"`
	GoalMonthsToBudget     *int32          `json:"goal_months_to_budget"`
	GoalUnderFunded        *Milliunits     `json:"goal_under_funded"`
	GoalOverallFunded      *Milliunits     `json:"goal_overall_funded"`
	GoalOverallLeft        *Milliunits     `json:"goal_overall_left"`
	Deleted                bool            `json:"deleted"`
}

type CategoryGroup struct {
	ID      CategoryGroupID `json:"id"`
	Name    string          `json:"name"`
	Hidden  bool            `json:"hidden"`
	Deleted bool            `json:"deleted"`
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

type DateFormat struct {
	Format string `json:"format"`
}

type LoadAccountPeriodicValue map[any]int64

type MonthDetail struct {
	Month        string     `json:"month"`
	Note         *string    `json:"note"`
	Income       Milliunits `json:"income"`
	Budgeted     Milliunits `json:"budgeted"`
	Activity     Milliunits `json:"activity"`
	ToBeBudgeted Milliunits `json:"to_be_budgeted"`
	AgeOfMoney   *int32     `json:"age_of_money"`
	Deleted      bool       `json:"deleted"`
	Categories   []Category `json:"categories"`
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

type PostAccountWrapper struct {
	Account SaveAccount `json:"account"`
}

type SaveAccount struct {
	Name    string      `json:"name"`
	Type    AccountType `json:"type"`
	Balance Milliunits  `json:""`
}

type ScheduledSubTransaction struct {
	ID                     ScheduledSubTransactionID `json:"id"`
	ScheduledTransactionID ScheduledTransactionID    `json:"scheduled_transaction_id"`
	Amount                 Milliunits                `json:"amount"`
	Memo                   *string                   `json:"memo"`
	PayeeID                *PayeeID                  `json:"payee_id"`
	CategoryID             *CategoryID               `json:"category_id"`
	TransferAccountID      *AccountID                `json:"transfer_account_id"`
	Deleted                bool                      `json:"deleted"`
}

type ScheduledTransactionSummary struct {
	ID                ScheduledTransactionID `json:"id"`
	DateFirst         string                 `json:"date_first"`
	DateNext          string                 `json:"date_next"`
	Frequency         ScheduleFrequency      `json:"frequency"`
	Amount            Milliunits             `json:"amount"`
	Memo              *string                `json:"memo"`
	FlagColor         FlagColor              `json:"flag_color"`
	AccountID         AccountID              `json:"account_id"`
	PayeeID           *PayeeID               `json:"payee_id"`
	CategoryID        *CategoryID            `json:"category_id"`
	TransferAccountID *AccountID             `json:"transfer_account_id"`
	Deleted           bool                   `json:"deleted"`
}

type SubTransaction struct {
	ID                    SubTransactionID `json:"id"`
	TransactionID         TransactionID    `json:"transaction_id"`
	Amount                Milliunits       `json:"amount"`
	Memo                  *string          `json:"memo"`
	PayeeID               *PayeeID         `json:"payee_id"`
	PayeeName             *string          `json:"payee_name"`
	CategoryID            *CategoryID      `json:"category_id"`
	CategoryName          *string          `json:"category_name"`
	TransferAccountID     *AccountID       `json:"transfer_account_id"`
	TransferTransactionID *TransactionID   `json:"transfer_transaction_id"`
	Deleted               bool             `json:"deleted"`
}

type TransactionSummary struct {
	ID                      TransactionID        `json:"id"`
	Date                    string               `json:"date"`
	Amount                  Milliunits           `json:"amount"`
	Memo                    *string              `json:"memo"`
	TransactionStatus       TransactionStatus    `json:"cleared"`
	Approved                bool                 `json:"approved"`
	FlagColor               FlagColor            `json:"flag_color"`
	AccountID               AccountID            `json:"account_id"`
	PayeeID                 *PayeeID             `json:"payee_id"`
	CategoryID              *CategoryID          `json:"category_id"`
	TransferAccountID       *AccountID           `json:"transfer_account_id"`
	TransferTransactionID   *TransactionID       `json:"transfer_transaction_id"`
	MatchedTransactionID    *TransactionID       `json:"matched_transaction_id"`
	ImportID                *ImportID            `json:"import_id"`
	ImportPayeeName         *string              `json:"import_payee_name"`
	ImportPayeeNameOriginal *string              `json:"import_payee_name_original"`
	DebtTransactionType     *DebtTransactionType `json:"debt_transaction_type"`
	Deleted                 bool                 `json:"deleted"`
}
