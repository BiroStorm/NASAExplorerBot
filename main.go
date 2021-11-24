package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	tele "gopkg.in/tucnak/telebot.v3"
)

const api = "https://api.nasa.gov/"

var bot *tele.Bot

func main() {

	// Configuration of env.
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w ", err))
	}
	fmt.Println("Configuration ok.")
	fmt.Println("Launching the bot...")

	b, err := tele.NewBot(tele.Settings{
		Token:  viper.GetString("BotToken"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot = b

	fmt.Printf("%v has been Launched! 🔭\n", bot.Me.Username)

	b.Handle("/start", func(m tele.Context) error {
		return m.Send("👨‍🚀 Welcome my fellow Astronaut 👩‍🚀\nUse /Today for the Picture or Video of the day!")
	})

	b.Handle("/Today", onToday)

	b.Start()
}
