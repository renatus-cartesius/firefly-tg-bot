package firefly

import (
	"time"
)

type AccountsResponse struct {
	Data AccountsData `json:"data"`
}

type AccountsData struct {
	Attributes Account `json:"attributes"`
	Type       string  `json:"type"`
	ID         string  `json:"id"`
}
type Account struct {
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
	Active                bool        `json:"active"`
	Order                 int         `json:"order"`
	Name                  string      `json:"name"`
	Type                  string      `json:"type"`
	AccountRole           string      `json:"account_role"`
	CurrencyID            string      `json:"currency_id"`
	CurrencyCode          string      `json:"currency_code"`
	CurrencySymbol        string      `json:"currency_symbol"`
	CurrencyDecimalPlaces int         `json:"currency_decimal_places"`
	CurrentBalance        string      `json:"current_balance"`
	CurrentBalanceDate    time.Time   `json:"current_balance_date"`
	Notes                 interface{} `json:"notes"`
	MonthlyPaymentDate    interface{} `json:"monthly_payment_date"`
	CreditCardType        interface{} `json:"credit_card_type"`
	AccountNumber         interface{} `json:"account_number"`
	Iban                  interface{} `json:"iban"`
	Bic                   interface{} `json:"bic"`
	VirtualBalance        string      `json:"virtual_balance"`
	OpeningBalance        string      `json:"opening_balance"`
	OpeningBalanceDate    interface{} `json:"opening_balance_date"`
	LiabilityType         interface{} `json:"liability_type"`
	LiabilityDirection    interface{} `json:"liability_direction"`
	Interest              interface{} `json:"interest"`
	InterestPeriod        interface{} `json:"interest_period"`
	CurrentDebt           interface{} `json:"current_debt"`
	IncludeNetWorth       bool        `json:"include_net_worth"`
	Longitude             interface{} `json:"longitude"`
	Latitude              interface{} `json:"latitude"`
	ZoomLevel             interface{} `json:"zoom_level"`
}

func (ar *AccountsResponse) String() string {
	res := "\xF0\x9F\x92\xB3 Аккаунт: " + ar.Data.Attributes.Name + "\n"
	res += "\xF0\x9F\x92\xB5 Баланс: " + ar.Data.Attributes.CurrentBalance + ar.Data.Attributes.CurrencySymbol + "\n"
	return res
}
