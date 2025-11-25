package handlers

import (
	"strings"

	"kafka/src/config"

	"github.com/bwmarrin/discordgo"
)

func Router(s *discordgo.Session, m *discordgo.MessageCreate, cfg *config.Config) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, cfg.Bot.Prefix) {
		return
	}

	content := strings.TrimPrefix(m.Content, cfg.Bot.Prefix)
	args := strings.Fields(content)

	if len(args) == 0 {
		return
	}

	command := strings.ToLower(args[0])

	switch command {
	case "ping":
		HandlePing(s, m)
	}
}
