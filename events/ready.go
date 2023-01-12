package ready

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mrtuxa/discordbotgo-boty/utils/command"
)

func Func(session *discordgo.Session, message *discordgo.MessageCreate) {
	command.IfBot(session, message)
	command.LCommand("ping", session, message, "Pong!")
}
