package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	bot, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))

	if err != nil {
		fmt.Println(err)
	}

	err = bot.Open()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bot is now running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}
