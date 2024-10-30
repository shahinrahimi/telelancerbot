package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) BlockBots(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		if u.Message.From == nil || u.Message.From.IsBot {
			b.l.Printf("Blocked bot: %d", u.Message.From.ID)
			return
		}
		next(u, ctx)
	}
}
