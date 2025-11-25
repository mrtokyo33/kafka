package handlers

import (
	"fmt"
	"kafka/src/services"

	"github.com/bwmarrin/discordgo"
)

func HandleMemeSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	options := i.ApplicationCommandData().Options
	allowNSFW := false

	for _, opt := range options {
		if opt.Name == "nsfw" {
			allowNSFW = opt.BoolValue()
		}
	}

	meme, err := services.GetRandomMeme("MemesBR")
	if err != nil {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &[]string{"❌ Failed to fetch meme from r/MemesBR. (API might be busy)"}[0],
		})
		return
	}

	if meme.NSFW && !allowNSFW {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &[]string{"⚠️ Fetched meme was NSFW. Retry with nsfw:True."}[0],
		})
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: meme.Title,
		URL:   meme.PostLink,
		Color: 0x00ff00,
		Image: &discordgo.MessageEmbedImage{URL: meme.URL},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("👍 %d | u/%s | r/%s", meme.Ups, meme.Author, meme.Subreddit),
		},
	}

	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{embed},
	})
}

func HandleMemeText(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	allowNSFW := false

	for _, arg := range args {
		if arg == "nsfw" {
			allowNSFW = true
		}
	}

	s.ChannelTyping(m.ChannelID)

	meme, err := services.GetRandomMeme("MemesBR")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed to fetch meme from r/MemesBR.")
		return
	}

	if meme.NSFW && !allowNSFW {
		s.ChannelMessageSend(m.ChannelID, "⚠️ The fetched meme was NSFW. Please use the `nsfw` flag.")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: meme.Title,
		URL:   meme.PostLink,
		Color: 0x00ff00,
		Image: &discordgo.MessageEmbedImage{URL: meme.URL},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("👍 %d | u/%s | r/%s", meme.Ups, meme.Author, meme.Subreddit),
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
