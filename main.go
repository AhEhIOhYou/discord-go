package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	usd := &discordgo.UpdateStatusData{
		Status: "online",
	}

	usd.Activities = []*discordgo.Activity{{
		Name: "with lolis",
		Type: discordgo.ActivityTypeGame,
	}}

	dg.UpdateStatusComplex(*usd)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "тест" {
		s.ChannelMessageSend(m.ChannelID, "test! - "+m.Author.Username+m.Author.ID)
		return
	}

	if m.Content == "тест!" {
		s.ChannelMessageSend(m.ChannelID, "мы твою мать всем отделом ебали")
		return
	}

	if m.Content == "test" {
		return
	}

	content := m.Content
	lookFor := "cat"
	contain := strings.Contains(content, lookFor)

	if contain {
		words := strings.Fields(content)
		if len(words) > 1 {
			s.ChannelMessageSendEmbed(m.ChannelID, getCatCodePict(words[1]))
		}
	}

}

func getCatCodePict(catCode string) *discordgo.MessageEmbed {

	emb := discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: "https://http.cat/" + catCode,
		},
	}

	return &emb
}
