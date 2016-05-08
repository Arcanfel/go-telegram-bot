
package telegram

import (
	"net/http"
	"net/url"
	"fmt"
)

const (
	defaultBaseURL = "https://api.telegram.org/bot"
	botToken = "YOUR_TOKEN_HERE"

	getMe = "getMe"
)
// A Client manages communication with Telegram API
type Client struct {
	*APIClient

	Updates *UpdatesService
	Bot 	*BotService
}

// Test gfgdgf
func (c *Client) Test() {
	println("abs")
}

// NewClient returns a new Telegram API client
func NewClient(httpClient *http.Client) *Client {

	u := fmt.Sprintf("%s%s/", defaultBaseURL, botToken)

	baseURL, _ := url.Parse(u)

	client := new(Client)
	client.APIClient 	= NewAPIClient(httpClient, baseURL)
	client.Updates 		= &UpdatesService{client: client}
	client.Bot 			= &BotService{client: client}

	return client
}