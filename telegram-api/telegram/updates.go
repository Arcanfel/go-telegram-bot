
package telegram

// UpdatesService handles communication with the updates for telegram
type UpdatesService struct {
	client *Client
}

// Update represents an update from telegram service
 type Update struct {
	 UpdateID 			string
	 Message 			string
	 InlineQuery 		string
	 ChoseInlineResult 	string
	 CallbackQuery 		string
 }