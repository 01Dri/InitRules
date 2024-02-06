package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var botID string
var tokenPrefix string = "Bot "                // Token for discord need Bot prefix
var guildID string = "1200125230203539569"     // ID server
var guildRoleId string = "1200133310987374664" // ID role main

func GetTokenBot() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}
	tokenValue := os.Getenv("TOKEN_DISCORD") // ENVIROMENT VARIABLE WITH TOKEN
	token := tokenPrefix + tokenValue
	return token
}
