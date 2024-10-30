package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/shahinrahimi/go-freelancer-sdk/v1"
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

	// list projects
	token2 := os.Getenv("FREELANCER_ACCESS_TOKEN")
	if token2 == "" {
		l.Fatal("FREELANCER_ACCESS_TOKEN is not set")
	}
	fc := freelancer.NewClient(token2)
	ps := fc.NewListActiveProjectsService()
	pr, err := ps.Do(ctx)
	if err != nil {
		l.Printf("Failed to list projects: %v", err)
		return
	}
	ownersIDs := make([]int, 0)

	for _, project := range pr.Result.Projects {
		l.Printf("Project: %s", project.Title)
		ownersIDs = append(ownersIDs, project.OwnerID)
	}
	us := fc.NewListUsersService()
	us.SetUsers(ownersIDs)
	ur, err := us.Do(ctx)
	if err != nil {
		l.Printf("Failed to list users: %v", err)
		return
	}
	for _, user := range ur.Result.Users {
		l.Printf("User: %s", user.Username)
	}

	// fc := freelancer.NewClient(token2)
	// projects, err := ps.Do(ctx)
	// if err != nil {
	// 	l.Printf("Failed to list projects: %v", err)
	// }

	// for _, project := range projects.Result.Projects {
	// 	l.Printf("Project: %s", project.Title)
	// }

	// us := fc.New

	cc := make(chan os.Signal, 1)
	signal.Notify(cc, os.Interrupt)
	<-cc
	time.AfterFunc(300*time.Millisecond, cancel)
	<-cc
}
