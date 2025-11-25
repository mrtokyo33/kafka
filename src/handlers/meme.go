package handlers

import (
	"fmt"
	"kafka/src/config"
	"kafka/src/services"

	"github.com/bwmarrin/discordgo"
)

func HandleMemeSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	options := i.ApplicationCommandData().Options
	subreddit := ""

	for _, opt := range options {
		if opt.Name == "subreddit" {
			subreddit = opt.StringValue()
		}
	}

	meme, err := services.GetMeme(subreddit)
	if err != nil {
		msg := "❌ Failed to fetch meme."
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &msg,
		})
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: meme.Title,
		URL:   meme.PostLink,
		Color: config.MemeColor,
		Image: &discordgo.MessageEmbedImage{URL: meme.URL},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("%s %d | u/%s | r/%s", config.MemeEmojiUpvote, meme.Ups, meme.Author, meme.Subreddit),
		},
	}

	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{embed},
	})
}

func HandleMemeText(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	subreddit := ""

	for _, arg := range args {
		subreddit = arg
	}

	s.ChannelTyping(m.ChannelID)

	meme, err := services.GetMeme(subreddit)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "❌ Failed to fetch meme.")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: meme.Title,
		URL:   meme.PostLink,
		Color: config.MemeColor,
		Image: &discordgo.MessageEmbedImage{URL: meme.URL},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("%s %d | u/%s | r/%s", config.MemeEmojiUpvote, meme.Ups, meme.Author, meme.Subreddit),
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
