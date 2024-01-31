package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	enviromentVariableToken := "TOKEN_DISCORD_INIT_RULES"
	var token string = os.Getenv(enviromentVariableToken)
	dg, err := discordgo.New(token)

	if err != nil {
		fmt.Println(err)
		return
	}
	// dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
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
