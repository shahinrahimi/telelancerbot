package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	l   *log.Logger
	api *tgbotapi.BotAPI
}

func NewBot(l *log.Logger, token string) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		l.Fatalf("Failed to create bot: %v", err)
	}
	return &Bot{
		l:   l,
		api: api,
	}
}

func (b *Bot) Start(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)
	b.l.Println("Bot started and listening for updates")
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case update := <-updates:
			if update.Message == nil {
				continue
			}
			b.l.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
	}
}

func (b *Bot) Shutdown() {
	b.l.Println("Shutting down ...")
	b.api.StopReceivingUpdates()
}
