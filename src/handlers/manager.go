package handlers

import "github.com/bwmarrin/discordgo"

type CommandDef struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var Commands = []CommandDef{
	{
		Definition: &discordgo.ApplicationCommand{
			Name:        "meme",
			Description: "Get a random meme from r/MemesBR",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "nsfw",
					Description: "Allow NSFW content",
					Required:    false,
				},
			},
		},
		Handler: HandleMemeSlash,
	},
}

func GetCommandMap() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	cmdMap := make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
	for _, cmd := range Commands {
		cmdMap[cmd.Definition.Name] = cmd.Handler
	}
	return cmdMap
}
