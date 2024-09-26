package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/shahinrahimi/telelancerbot/bot"
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

	go b.Start(ctx)

	cc := make(chan os.Signal, 1)
	signal.Notify(cc, os.Interrupt)
	<-cc

	b.Shutdown()

}
