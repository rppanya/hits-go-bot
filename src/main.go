package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotState struct {
	State map[int64]string
}

func NewBotState() *BotState {
	return &BotState{State: make(map[int64]string)}
}

func (b *BotState) GetState(chatID int64) string {
	return b.State[chatID]
}

func (b *BotState) SetState(chatID int64, state string) {
	b.State[chatID] = state
}

func main() {
	botToken := "7695784459:AAEVmeYca6RzvW6c5NIrqiOpGbTg9viH-_w"

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)
	state := NewBotState()

	for update := range updates {
		if update.Message == nil { // Ignore any non-Message Updates
			continue
		}

		chatID := update.Message.Chat.ID
		currentState := state.GetState(chatID)

		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(chatID, "Welcome! Please type something to start the conversation.")
			bot.Send(msg)
			state.SetState(chatID, "awaiting_input")
		} else if currentState == "awaiting_input" {
			response := fmt.Sprintf("You said: %s", update.Message.Text)
			msg := tgbotapi.NewMessage(chatID, response)
			bot.Send(msg)
			state.SetState(chatID, "input_received")
		} else {
			msg := tgbotapi.NewMessage(chatID, "I am not sure what to do. Use /start to restart.")
			bot.Send(msg)
		}
	}
}
