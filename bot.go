package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func StartBot() {

	dg, err := discordgo.New(GetTokenBot())
	if err != nil {
		fmt.Println(err)
		return
	}
	dg.AddHandlerOnce(OnReady)
	dg.AddHandler(MessageCreate)
	dg.AddHandler(OnJoinSetDefaultRole)

	dg.Identify.Intents = discordgo.IntentGuildMessages
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}
	fmt.Println("initRules is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	defer dg.Close()
}

func OnReady(bot *discordgo.Session, event *discordgo.Ready) {
	botID = bot.State.User.ID
	fmt.Println("Bot is ready. Bot ID:", botID)
}
