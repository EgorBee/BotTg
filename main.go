package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

func main() {
	apiurl := "https://api.telegram.org/bot"
	apitoken := ""
	url := apiurl + apitoken
	offset := 0

	for {
		updates, err := getUpdates(url, offset)
		if err != nil {
			log.Println("Something went wrong: ", err.Error())
		}
		for _, update := range updates {
			respond(url, update)
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)

	}

}

func getUpdates(url string, offset int) ([]Update, error) {
	resp, err := http.Get(url + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(url string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId

	if update.Message.Text == "/start" {
		// В чат вошел новый пользователь
		// Поприветствуем его
		go start()
		go start1()

		botMessage.Text = "Приветствие"
		buf, err := json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		return nil
	} else if update.Message.Text == "Расскажи что-нибудь" {
		botMessage.Text = random()
		buf, err := json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		return nil
	}
	return nil

}

func SendMessage() error {
	go startD()
	apiurl := "https://api.telegram.org/bot"
	apitoken := ""
	url := apiurl + apitoken
	var botMessage BotMessage
	botMessage.ChatId = 0 //id пользователя 1
	botMessage.Text = morning()
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil

}

func SendMessage2() error {
	apiurl := "https://api.telegram.org/bot"
	apitoken := ""
	url := apiurl + apitoken
	var botMessage BotMessage
	botMessage.ChatId = 00 //id пользователя 2
	botMessage.Text = morning()
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil

}

func SendMessageDAY() error {
	apiurl := "https://api.telegram.org/bot"
	apitoken := ""
	url := apiurl + apitoken
	var botMessage BotMessage
	botMessage.ChatId = 0 //id пользователя 1
	botMessage.Text = day()
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func SendMessage2DAY() error {
	apiurl := "https://api.telegram.org/bot"
	apitoken := ""
	url := apiurl + apitoken
	var botMessage BotMessage
	botMessage.ChatId = 00 //id пользователя 2
	botMessage.Text = day()
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil

}

func SendMessageEVENING() error {
	apiurl := "https://api.telegram.org/bot"
	apitoken := ""
	url := apiurl + apitoken
	var botMessage BotMessage
	botMessage.ChatId = 0 //id пользователя 1
	botMessage.Text = evening()
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func SendMessage2EVENING() error {
	apiurl := "https://api.telegram.org/bot"
	apitoken := ""
	url := apiurl + apitoken
	var botMessage BotMessage
	botMessage.ChatId = 00 //id пользователя 2
	botMessage.Text = evening()
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil

}

func start() error {
	apiurl := "https://api.telegram.org/bot"
	apitoken := ""
	url := apiurl + apitoken
	var botMessage BotMessage
	botMessage.ChatId = 00 //id пользователя 2
	botMessage.Text = "Началось"
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func startM() {
	timer1 := time.NewTicker(24 * time.Hour)

	for i := 0; i < 3650; i++ {
		select {
		case <-timer1.C:
			go SendMessage()
			go SendMessage2()

		}
	}
}

func start1() {

	timer1 := time.NewTicker(12 * time.Hour) // отложенный старт

	select {
	case <-timer1.C:
		go startM()
		SendMessage()
		SendMessage2()
	}
}

func startD() {
	timer1 := time.NewTicker(6 * time.Hour) // отложенный старт

	select {
	case <-timer1.C:
		go startE()
		SendMessageDAY()
		SendMessage2DAY()

	}

}

func startE() {
	timer1 := time.NewTicker(7 * time.Hour) // отложенный старт

	select {
	case <-timer1.C:
		SendMessageEVENING()
		SendMessage2EVENING()

	}

}

func morning() string {

	rand.Seed(time.Now().UnixNano())
	var messages [2]string
	var output string
	messages[0] = "Доброе утро!"
	messages[1] = "Утро доброе!"

	numberOfMessages := (rand.Intn(2))
	output = messages[numberOfMessages]
	return output
}

func random() string {

	rand.Seed(time.Now().UnixNano())
	var messages [2]string
	var output string
	messages[0] = "Фраза 1"
	messages[1] = "Фраза 2"

	numberOfMessages := (rand.Intn(2))
	output = messages[numberOfMessages]
	return output
}

func day() string {

	rand.Seed(time.Now().UnixNano())
	var messages [2]string
	var output string
	messages[0] = "Добрый день!"
	messages[1] = "День добрый!"

	numberOfMessages := (rand.Intn(2))
	output = messages[numberOfMessages]
	return output
}

func evening() string {

	rand.Seed(time.Now().UnixNano())
	var messages [2]string
	var output string
	messages[0] = "Добрый вечер!"
	messages[1] = "Вечер добрый!"

	numberOfMessages := (rand.Intn(2))
	output = messages[numberOfMessages]
	return output
}
