package firefly

import (
	"encoding/json"
	"firebot/internal/utils"
	"fmt"
	"log"
)

type FireflyClient struct {
	url        string
	token      string
	authClient *utils.AuthClient
}

func NewFireflyClient(url string, token string) *FireflyClient {
	return &FireflyClient{
		url:        url,
		token:      token,
		authClient: utils.NewAuthClient(token),
	}
}

func (c *FireflyClient) Ping() {
	fmt.Println("Ping to firefly server")
}

func (c *FireflyClient) GetTransactions() (string, error) {
	return c.authClient.Get(c.url + "/transactions")
}

func (c *FireflyClient) GetAccounts() (string, error) {
	data, err := c.authClient.Get(c.url + "/accounts")
	if err != nil {
		log.Fatalln("Error on making request for Account: err")
	}

	res := &AccountsResponse{}
	if err = json.Unmarshal([]byte(data), &res); err != nil {
		log.Fatalln("Error on unmarshaling json Account:", err)
	}
	return res.String(), nil
}

func (c *FireflyClient) GetBudget() (string, error) {
	data, err := c.authClient.Get(c.url + "/budgets?start=2024-02-01&end=2024-02-29") // TODO: fix to using current month interval after tests
	if err != nil {
		log.Fatalln("Error on making request for Account: err")
	}

	res := &BudgetResponse{}
	if err = json.Unmarshal([]byte(data), &res); err != nil {
		log.Fatalln("Error on unmarshaling json Account:", err)
	}
	return res.String(), nil
}
