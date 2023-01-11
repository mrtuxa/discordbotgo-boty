package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	ready "github.com/mrtuxa/discordbotgo-boty/events"
	"github.com/mrtuxa/discordbotgo-boty/utils/merror"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {

	}
	merror.CheckFatal(err, "GoDotenv", "Can't load environment")

	envToken := os.Getenv("TOKEN")

	if envToken == "" {
		os.Exit(1)
	}
	token := "Bot" + envToken

	client, err := discordgo.New(token)
	merror.DisGoClientError(err)

	client.AddHandler(ready.ReadyFunc)
	client.Identify.Intents = discordgo.IntentsAll

	err = client.Open()
	merror.CheckPrintReturn(err, "[DiscordGo] Client:", "error opening connection", err)

	fmt.Println("Bot is now running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = client.Close()
	merror.CheckReturn(err)
}
