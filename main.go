package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var botID string

func main() {
	enviromentVariableToken := "TOKEN_DISCORD_INIT_RULES"
	var token string = os.Getenv(enviromentVariableToken)
	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	dg.AddHandlerOnce(onReady)
	dg.Identify.Intents = discordgo.IntentGuildMessages
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}
	fmt.Println("initRules is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()

}
func onReady(bot *discordgo.Session, event *discordgo.Ready) {
	botID = bot.State.User.ID
	fmt.Println("Bot is ready. Bot ID:", botID)
}

func messageCreate(bot *discordgo.Session, message *discordgo.MessageCreate) {
	fmt.Printf("Author: %s (%s)\n", message.Author.Username, message.Author.ID)
	fmt.Printf("Content: %s\n", message.Content)

	if message.Author.ID == botID {
		return
	}

	if message.Content == "!start" {
		fmt.Println(message.Content)
		_, err := bot.ChannelMessageSend(message.ChannelID, "hello, world!!")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("nao ta lendo a mensagem n dog")
	}
}
