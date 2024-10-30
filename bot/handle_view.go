package bot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleView(update *tgbotapi.Update, ctx context.Context) error {
	return nil
}

func (b *Bot) HandleViewCountries(update *tgbotapi.Update, ctx context.Context) error {
	countries, err := b.fc.GetCountries(ctx)
	if err != nil {
		return fmt.Errorf("failed to get countries: %v", err)
	}
	msg := ""
	for index, country := range countries {
		msg = msg + fmt.Sprintf("%d ", index) + country.Name + " " + country.Code + "\n"
	}
	b.MsgChan <- BotMessage{
		ChatID: update.Message.From.ID,
		Msg:    msg,
	}
	return nil
}
