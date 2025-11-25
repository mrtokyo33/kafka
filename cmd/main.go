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

	err = sess.Open()
	if err != nil {
		log.Fatalf("%sError opening connection: %v%s", config.ColorRed, err, config.ColorReset)
	}

	fmt.Printf("%s%s v%s is running...%s\n", config.ColorCyan, config.AppName, config.Version, config.ColorReset)
	fmt.Printf("%sPrefix: %s%s\n", config.ColorYellow, cfg.Bot.Prefix, config.ColorReset)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	sess.Close()
	fmt.Printf("\n%sBot shutdown successfully.%s\n\n", config.ColorRed, config.ColorReset)
}
