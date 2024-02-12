package firefly

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type BudgetResponse struct {
	Data []BudgetData `json:"data"`
}

type BudgetData struct {
	Attributes Budget `json:"attributes"`
}

type Budget struct {
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
	Active                 bool        `json:"active"`
	Name                   string      `json:"name"`
	Order                  int         `json:"order"`
	Notes                  interface{} `json:"notes"`
	AutoBudgetType         string      `json:"auto_budget_type"`
	AutoBudgetPeriod       string      `json:"auto_budget_period"`
	AutoBudgetCurrencyID   string      `json:"auto_budget_currency_id"`
	AutoBudgetCurrencyCode string      `json:"auto_budget_currency_code"`
	AutoBudgetAmount       string      `json:"auto_budget_amount"`
	Spent                  []struct {
		Sum                   string `json:"sum"`
		CurrencyID            int    `json:"currency_id"`
		CurrencyName          string `json:"currency_name"`
		CurrencySymbol        string `json:"currency_symbol"`
		CurrencyCode          string `json:"currency_code"`
		CurrencyDecimalPlaces int    `json:"currency_decimal_places"`
	} `json:"spent"`
}

func (br *BudgetResponse) String() (res string) {
	for _, budget := range br.Data {
		log.Println(budget)
		amount, err := strconv.ParseFloat(budget.Attributes.AutoBudgetAmount, 64)
		if err != nil {
			continue
		}
		spent, err := strconv.ParseFloat(budget.Attributes.Spent[0].Sum, 64)
		if err != nil {
			continue
		}
		res += "\xF0\x9F\x93\x8C Бюджет: " + budget.Attributes.Name + "\n"
		res += fmt.Sprintf("\xF0\x9F\x93\x8A Лимит: `%f ₽`\n", amount)
		res += fmt.Sprintf("\xF0\x9F\x93\x8A Осталось: `%f ₽`\n", amount+spent)
		res += "\n"
	}
	return res
}
