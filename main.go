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

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "тест" {
		s.ChannelMessageSend(m.ChannelID, "test! - "+m.Author.Username+m.Author.ID)
	}

	if m.Content == "тест!" {
		s.ChannelMessageSend(m.ChannelID, "мы твою мать всем отделом ебали")
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
