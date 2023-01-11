package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mrtuxa/discordbotgo-boty/utils/merror"
)

func IfBot(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
}

func SendMessage(session *discordgo.Session, m *discordgo.MessageCreate, message string) (*discordgo.Message, error) {
	return session.ChannelMessage(m.ChannelID, message)
}

func LCommand(command string, s *discordgo.Session, m *discordgo.MessageCreate, response string) {
	if m.Content == ">"+command {
		_, err := SendMessage(s, m, response)
		merror.CheckReturn(err)
	}
}
