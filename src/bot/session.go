package bot

import (
	"github.com/bwmarrin/discordgo"
)

func NewSession(token string, intentNames []string) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	var intents coreIntent
	for _, name := range intentNames {
		switch name {
		case "GUILD_MESSAGES":
			intents |= discordgo.IntentsGuildMessages
		case "MESSAGE_CONTENT":
			intents |= discordgo.IntentsMessageContent
		case "GUILDS":
			intents |= discordgo.IntentsGuilds
		}
	}

	session.Identify.Intents = discordgo.MakeIntent(intents)

	return session, nil
}

type coreIntent = discordgo.Intent
