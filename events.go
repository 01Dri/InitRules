package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == botID {
		return // Ignore the message of bot
	}

	if message.Content == "!start" {
		fmt.Println(message.Content)
		_, err := bot.ChannelMessageSend(message.ChannelID, "hello, world!!")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func OnJoinSetDefaultRole(session *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user != nil {
		if message.Type == discordgo.MessageTypeGuildMemberJoin {
			user := message.Author
			fmt.Println("Usuário entrou no servidor:", user.Username)
			session.GuildMemberRoleAdd(guildID, user.ID, guildRoleId)
			fmt.Println("Cargo main setado! para o usuario: ", user.Username)
		}
	}
}
