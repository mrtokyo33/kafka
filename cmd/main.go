package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"kafka/src/bot"
	"kafka/src/config"
	"kafka/src/handlers"

	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("%sError loading config: %v%s", config.ColorRed, err, config.ColorReset)
	}

	sess, err := bot.NewSession(cfg.Bot.Token, cfg.Bot.Intents)
	if err != nil {
		log.Fatalf("%sError creating session: %v%s", config.ColorRed, err, config.ColorReset)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.Router(s, m, cfg)
	})

	sess.AddHandler(handlers.SlashRouter)

	err = sess.Open()
	if err != nil {
		log.Fatalf("%sError opening connection: %v%s", config.ColorRed, err, config.ColorReset)
	}

	fmt.Printf("%s%s v%s is running...%s\n", config.ColorCyan, config.AppName, config.Version, config.ColorReset)
	fmt.Printf("%sPrefix: %s%s\n", config.ColorYellow, cfg.Bot.Prefix, config.ColorReset)

	if cfg.Bot.GuildID != "" {
		log.Printf("%sRegistering commands to Guild ID: %s (Instant Update)%s", config.ColorBlue, cfg.Bot.GuildID, config.ColorReset)
	} else {
		log.Printf("%sRegistering Global Commands (Up to 1 hour delay)%s", config.ColorBlue, config.ColorReset)
	}

	var commands []*discordgo.ApplicationCommand
	for _, v := range handlers.Commands {
		commands = append(commands, v.Definition)
	}

	_, err = sess.ApplicationCommandBulkOverwrite(sess.State.User.ID, cfg.Bot.GuildID, commands)
	if err != nil {
		log.Printf("%sError overwriting commands: %v%s", config.ColorRed, err, config.ColorReset)
	} else {
		log.Printf("%sCommands successfully updated!%s", config.ColorGreen, config.ColorReset)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	sess.Close()
	fmt.Printf("\n%sBot shutdown successfully.%s\n", config.ColorRed, config.ColorReset)
}
