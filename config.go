package main

import "os"

var botID string
var tokenPrefix string = "Bot "                               // Token for discord need Bot prefix
var guildID string = "1200125230203539569"                    // ID server
var guildRoleId string = "1200133310987374664"                // ID role main
var tokenValue string = os.Getenv("TOKEN_DISCORD_INIT_RULES") // ENVIROMENT VARIABLE WITH TOKEN

func GetTokenBot() string {
	token := tokenPrefix + tokenValue
	return token
}
