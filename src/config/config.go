package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Bot struct {
		Token   string
		GuildID string
		Prefix  string
		Intents []string
	}
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	cfg.Bot.Token = os.Getenv("DISCORD_TOKEN")
	if cfg.Bot.Token == "" {
		return nil, fmt.Errorf("DISCORD_TOKEN not found in .env")
	}

	cfg.Bot.GuildID = os.Getenv("DISCORD_GUILD_ID")

	return &cfg, nil
}
