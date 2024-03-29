package botUtils

import (
	"log"
	"net/http"
	"strconv"

	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/utils"
)

// SendTextMessage 待修改
func SendTextMessage(chatId int, data string) {
	config := globals.UnmarshaledConfig

	url := config.Bot.APIUrl + config.Bot.Token + config.Bot.Methods.SendMessage
	params := map[string]string{
		"chat_id":    strconv.Itoa(chatId),
		"text":       data,
		"parse_mode": "MarkdownV2", // Send as Markdown text
	}
	code, body, err := utils.HttpGET(url, params)
	if err != nil || code != http.StatusOK {
		log.Println("Error sending TEXT message:", err)
		log.Println("Response Body:", string(body))
		return
	}
}

/*
func SendPhotoMessage(msg string) {
	message := handleMsg(msg)
	url := globals.UnmarshaledConfig.Bot.Token + "sendPhoto"
	params := map[string]string{
		"chat_id": message.ChatId,
		"photo":   message.Photo,
	}
	reqURL := buildURL(url, params)
	response, err := http.Get(reqURL)
	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing :", err)
		}
	}(response.Body)
	var buf bytes.Buffer
	_, err = io.Copy(&buf, response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	log.Println("Response Body:", buf.String())
}
*/

/*
func handleMsg(msg string) *models.Message {
	decodedMsgBytes, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		log.Println("Error decode base64:", err)
		return nil
	}
	msg = string(decodedMsgBytes)
	var message = new(models.Message)
	err = xml.Unmarshal([]byte(msg), &message)
	if err != nil {
		log.Println("Error unmarshalling XML:", err)
		return nil
	}
	return message
}
*/
