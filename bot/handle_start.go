package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleStart(update *tgbotapi.Update, ctx context.Context) error {
	return nil
}
