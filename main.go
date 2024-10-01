package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/shahinrahimi/telelancerbot/bot"
	"github.com/shahinrahimi/telelancerbot/types"
)

func main() {
	l := log.New(os.Stdout, "[TELELANCERBOT] ", log.LstdFlags)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := godotenv.Load(); err != nil {
		l.Fatalf("Failed to load .env file: %v", err)
	}
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		l.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}
	b := bot.NewBot(l, token)

	r := b.GetRouter()
	r.Use(b.Logger)

	vr := r.NewRoute("View")
	vr.HandleCommand(types.CommandView, b.MakeHandlerFunc(b.HandleView))

	go func() {
		if err := b.Start(ctx); err != nil {
			l.Fatal(err)
		}
	}()

	cc := make(chan os.Signal, 1)
	signal.Notify(cc, os.Interrupt)
	<-cc
	time.AfterFunc(300*time.Millisecond, cancel)
	<-cc
}
