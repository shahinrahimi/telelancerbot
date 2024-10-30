package bot

import (
	"context"
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/telelancerbot/models"
)

func (b *Bot) RequireAuthentication(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		user := ctx.Value(models.KeyUser{}).(*models.User)
		if user == nil {
			b.l.Panicf("user is nil")
		}
		if !user.IsConfirmed {
			b.l.Printf("user %d is not confirmed", user.ID)
			return
		}
		b.l.Printf("User identified: %d", user.ID)
		next(u, ctx)
	}
}

func (b *Bot) RequireAuthorization(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		user := ctx.Value(models.KeyUser{}).(*models.User)
		if user == nil {
			b.l.Panicf("user is nil")
		}
		if !user.IsAdmin {
			b.l.Printf("user %d is forbidden", user.ID)
			return
		}
		next(u, ctx)
	}
}

func (b *Bot) ProvideNewUser(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		userID := u.Message.From.ID
		chatID := u.Message.Chat.ID
		user, err := b.s.GetUser(userID)
		if err != nil {
			if err != sql.ErrNoRows {
				b.l.Printf("failed to get user %d from DB: %v", userID, err)
				return
			}
			b.l.Printf("User not found in DB userID: %d", userID)
			// try creating a user
			user = &models.User{
				ID:          userID,
				IsAdmin:     false,
				IsConfirmed: false,
			}
			if err := b.s.InsertUser(user); err != nil {
				b.l.Printf("unexpected error to insert user %d to DB: %v", userID, err)
				return
			}
			b.l.Printf("New user created: %d, from chat %d", userID, chatID)
			// try refetching user
			user, err = b.s.GetUser(userID)
			if err != nil {
				b.l.Printf("unexpected error to get user %d from DB: %v", userID, err)
				return
			}

		} else {
			b.l.Printf("User identified: %d", user.ID)
		}
		ctx = context.WithValue(ctx, models.KeyUser{}, *user)
		next(u, ctx)
	}
}
