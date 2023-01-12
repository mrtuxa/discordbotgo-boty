package main

import (
	"fmt"
	"github.com/mrtuxa/discordbotgo-boty/utils/webinterface"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	ready "github.com/mrtuxa/discordbotgo-boty/events"
	"github.com/mrtuxa/discordbotgo-boty/utils/merror"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	merror.CheckFatal(err, "GoDotenv", "Can't load environment")

	envToken := os.Getenv("TOKEN")
	portWebinterface := os.Getenv("WEBINTERFACE_PORT")

	if envToken == "" {
		os.Exit(1)
	}

	fmt.Println("Port")
	fmt.Println(portWebinterface)

	if len(portWebinterface) == 0 {
		fmt.Println("using Default Port 2103")
		portWebinterface = "2103"
	}

	webinterface.Start(portWebinterface)

	token := "Bot " + envToken

	client, err := discordgo.New(token)
	merror.DisGoClientError(err)

	client.AddHandler(ready.Func)
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
