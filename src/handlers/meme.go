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
	allowNSFW := false
	subreddit := ""

	for _, opt := range options {
		if opt.Name == "nsfw" {
			allowNSFW = opt.BoolValue()
		}
		if opt.Name == "subreddit" {
			subreddit = opt.StringValue()
		}
	}

	channel, err := s.Channel(i.ChannelID)
	if err != nil {
		return
	}

	if allowNSFW && !channel.NSFW {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &[]string{"❌ You cannot use `nsfw:True` in a non-NSFW channel."}[0],
		})
		return
	}

	meme, err := services.GetMeme(subreddit, allowNSFW)
	if err != nil {
		msg := "❌ Failed to fetch meme."
		if allowNSFW {
			msg += " (Could not find an NSFW meme in this subreddit after retries)"
		}
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
	allowNSFW := false
	subreddit := ""

	for _, arg := range args {
		if arg == "nsfw" {
			allowNSFW = true
		} else {
			subreddit = arg
		}
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return
	}

	if allowNSFW && !channel.NSFW {
		s.ChannelMessageSend(m.ChannelID, "❌ You cannot use the NSFW flag in a non-NSFW channel.")
		return
	}

	s.ChannelTyping(m.ChannelID)

	meme, err := services.GetMeme(subreddit, allowNSFW)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "❌ Failed to fetch meme (or couldn't find NSFW content).")
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
