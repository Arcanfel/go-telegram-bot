package telegram

// BotService handles communication of bot interaction with telegram service
type BotService struct {
	client *Client
}

// Bot represents a bot information from telegram service
 type Bot struct {
	 ID 		int `json:"id"`
	 FirstName 	string `json:"first_name"`
	 LastName 	string `json:"last_name"`
	 Username 	string `json:"username"`
 }

// Get retrieves information about the telegram bot the token belongs to
 func (s* BotService) Get() (*Bot, *Response, error) {

	 u, err := addOptions(getMe, nil)
	 if err != nil {
		 return nil, nil, err
	 }

	 req, err := s.client.NewRequest("GET", u, nil)
	 if err != nil {
		 return nil, nil, err
	 }

	 bot := new(Bot)

	 resp, err := s.client.Do(req, bot)
	 if err != nil {
		 return nil, resp, err
	 }

	 return bot, resp, nil
 }