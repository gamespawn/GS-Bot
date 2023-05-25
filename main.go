package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		fmt.Println(envErr)
		log.Fatalf("Error loading .env file")
	}

	bot, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))

	if err != nil {
		fmt.Print(err)
		log.Fatalf("Error making bot session")
	}

	err = bot.Open()

	if err != nil {
		fmt.Print(err)
		log.Fatalf("Error opening bot: ")
	}
	fmt.Println("Bot is now running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}
