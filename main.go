package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var botID string
var tokenPrefix string = "Bot "                // Token for discord need Bot prefix
var guildID string = "1179517927679078521"     // ID server
var guildRoleId string = "1202703396981964871" // ID role main

func getTokenBot() string {
	enviromentVariableToken := "TOKEN_DISCORD_INIT_RULES"
	var tokenValue string = os.Getenv(enviromentVariableToken)
	token := tokenPrefix + tokenValue
	return token
}
func main() {
	dg, err := discordgo.New(getTokenBot())
	if err != nil {
		fmt.Println(err)
		return
	}
	dg.AddHandlerOnce(onReady)
	dg.AddHandler(messageCreate)
	dg.AddHandler(onJoinSetDefaultRole)

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
func onReady(bot *discordgo.Session, event *discordgo.Ready) {
	botID = bot.State.User.ID
	fmt.Println("Bot is ready. Bot ID:", botID)
}

func messageCreate(bot *discordgo.Session, message *discordgo.MessageCreate) {
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

func onJoinSetDefaultRole(session *discordgo.Session, message *discordgo.MessageCreate) {
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
