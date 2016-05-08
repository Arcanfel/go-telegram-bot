package main

import (
	"fmt"
	"daed.net/telegram-api/telegram"
)

func main() {
	client := telegram.NewClient(nil)

	bot, _, err := client.Bot.Get()
	if err != nil {
		println(err.Error())
	}

	fmt.Printf("bot name is %s. hello\n", bot.FirstName)

	NewServer().Start()
}