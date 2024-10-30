package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/telelancerbot/models"
)

func (b *Bot) RequireAuthentication(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		userID := u.Message.From.ID
		chatID := u.Message.Chat.ID
		user, err := b.s.GetUser(userID)
		if err != nil {
			b.l.Printf("failed to get user %d from DB: %v", userID, err)
			return
		}
		b.l.Printf("User identified: %d, from chat %d", userID, chatID)
		ctx = context.WithValue(ctx, models.KeyUser{}, *user)
		next(u, ctx)
	}
}

func (b *Bot) RequireAuthorization(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		userID := u.Message.From.ID
		chatID := u.Message.Chat.ID
		user, err := b.s.GetUser(userID)
		if err != nil {
			b.l.Printf("failed to get user %d from DB: %v", userID, err)
			return
		}
		if !user.IsAdmin {
			b.l.Printf("user %d is forbidden", userID)
			return
		}
		b.l.Printf("User identified: %d, from chat %d", userID, chatID)
		ctx = context.WithValue(ctx, models.KeyUser{}, *user)
		next(u, ctx)
	}
}
