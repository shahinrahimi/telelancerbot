package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/telelancerbot/client"
	"github.com/shahinrahimi/telelancerbot/store"
	"github.com/shahinrahimi/telelancerbot/types"
)

type Bot struct {
	l       *log.Logger
	api     *tgbotapi.BotAPI
	router  *Router
	s       store.Storage
	fc      *client.Freelancer
	MsgChan chan BotMessage
}

func New(l *log.Logger, token string, s store.Storage, fc *client.Freelancer) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		l.Fatalf("Failed to create bot: %v", err)
	}
	return &Bot{
		l:       l,
		api:     api,
		fc:      fc,
		router:  newRouter(),
		s:       s,
		MsgChan: make(chan BotMessage),
	}
}

func newRouter() *Router {
	return &Router{
		middlewares: make([]Middleware, 0),
		handlers:    make(map[types.CommandType]Handler),
		routes:      make(map[string]*Route),
	}
}

func (b *Bot) GetRouter() *Router {
	return b.router
}

func (b *Bot) Start(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)
	b.l.Println("Bot started and listening for updates")
	for {
		select {
		case <-ctx.Done():
			b.api.StopReceivingUpdates()
			b.l.Println("Shutting down ...")
			return ctx.Err()
		case update := <-updates:
			if update.Message == nil {
				continue
			}
			b.handleUpdate(ctx, update)
		case msg := <-b.MsgChan:
			b.sendMessage(msg.ChatID, msg.Msg)
		}
	}
}

func (b *Bot) sendMessage(chatID int64, msgStr string) {
	msg := tgbotapi.NewMessage(chatID, msgStr)
	if _, err := b.api.Send(msg); err != nil {
		b.l.Printf("error in sending message to user: %v", err)
	}
}

func (b *Bot) MakeHandlerFunc(f ErrorHandler) Handler {
	return func(update *tgbotapi.Update, ctx context.Context) {
		if err := f(update, ctx); err != nil {
			b.l.Panic(err)
		}
	}
}

func (b *Bot) handleUpdate(ctx context.Context, u tgbotapi.Update) {
	c := u.Message.Command()
	command, err := types.StringToCommandType(c)
	if err != nil {
		// message the user the command not known
		return
	}

	finalHandler := func(u *tgbotapi.Update, ctx context.Context) {
		for _, route := range b.router.routes {
			if handler, exists := route.handlers[command]; exists {
				routeHandler := handler
				// apply route middlewares in reverse order
				for i := len(route.middlewares) - 1; i >= 0; i-- {
					routeHandler = route.middlewares[i](routeHandler)
				}
				routeHandler(u, ctx)
			}
		}
	}

	// apply global middlewares in reverse order
	for i := len(b.router.middlewares) - 1; i >= 0; i-- {
		finalHandler = b.router.middlewares[i](finalHandler)
	}
	finalHandler(&u, ctx)
}
