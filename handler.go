package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	tele "gopkg.in/tucnak/telebot.v3"
)

func onToday(c tele.Context) error {
	response, err := http.Get(api + "planetary/apod?api_key=DEMO_KEY")

	if err != nil {
		return c.Send("API Problem! 💔🚀")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var todayPicture APOD
	json.Unmarshal(responseData, &todayPicture)

	switch todayPicture.Type {
	case "image":
		photo := &tele.Photo{File: tele.FromURL(todayPicture.URL)}

		photo.Caption = formattingText(todayPicture.Title, todayPicture.Text, todayPicture.Copyright, func() bool {
			if strings.Compare(todayPicture.Copyright, "") == 0 {
				return false
			}
			return true
		}())
		_, err := bot.Send(c.Chat(), photo, tele.ModeHTML)
		return err
	case "video":
		message := todayPicture.URL + "\n\n"
		message += formattingText(todayPicture.Title, todayPicture.Text, todayPicture.Copyright, func() bool {
			if strings.Compare(todayPicture.Copyright, "") == 0 {
				return false
			}
			return true
		}())
		_, err := bot.Send(c.Chat(), message, tele.ModeHTML)
		return err
	}

	return nil
}

func formattingText(title, text, author string, copyrighted bool) string {
	formattedText := "<b>" + title + "</b>\n\n"

	bodyText := strings.ReplaceAll(strings.ReplaceAll(text, ". ", ".\n"), "\n ", "\n")

	// the limit for the caption on Photos are 1024, using lower value because of Copyright
	if strings.Count(formattedText+bodyText, "") > 1000 {
		substr := strings.Split(bodyText, "\n")
		for _, val := range substr {
			if strings.Count(formattedText+"\n"+val, "") > 1000 {
				break
			} else {
				formattedText += val + "\n"
			}
		}
	} else {
		formattedText += bodyText + "\n"
	}
	if copyrighted {
		formattedText += "\n<i>© " + author + "</i>"
	}
	return formattedText
}
