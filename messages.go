package telegramBot

import (
    "log"
    "strings"
    _ "strconv"
    _ "net/url"
    "net/http"
    "encoding/json"
)

// Sends a message with the given parameters
func (bot Bot) SendMessageWithParameters(message MessageRequest) Message {
    // Create the request
    requestURL := bot.baseURL + "/sendMessage"


    // Make the request
    response, err := http.Post(requestURL, "application/json", strings.NewReader(message.AsJSON()))
    if err != nil {
        log.Fatal(err)
    }

    // Parse the response
    var responseBody MessageResponse
    json.NewDecoder(response.Body).Decode(&responseBody)

    return *responseBody.Message
}

// Sends a text message to the given chat
func (bot Bot) SendMessage(body string, chatId int, parseMode string) Message {
    // Define the parameters
    var message MessageRequest
    message.ChatId = chatId
    message.Body = body
    message.ParseMode = parseMode

    return bot.SendMessageWithParameters(message)
}

// Sends a text message to a given chat with an inline keyboard
func (bot Bot) SendMessageWithKeyboard(body string, chatId int, parseMode string, replyMarkup InlineKeyboardMarkup) Message {
    // Define the parameters
    var message MessageRequest
    message.ChatId = chatId
    message.Body = body
    message.ParseMode = parseMode
    message.ReplyMarkup = &replyMarkup

    return bot.SendMessageWithParameters(message)
}

// Sends a Markdown message to the given chat
func (bot Bot) SendMarkdownMessage(body string, chatId int) Message {
    return bot.SendMessage(body, chatId, "Markdown")
}

// Sends an HTML message to the given chat
func (bot Bot) SendHTMLMessage(body string, chatId int) Message {
    return bot.SendMessage(body, chatId, "HTML")
}

func (bot Bot) EditMessageText(chatId int, messageId int, body string, parseMode string) {
    // Create the request
    requestURL := bot.baseURL + "/editMessageText"
    var parameters EditMessageTextRequest
    parameters.ChatId = chatId
    parameters.MessageId = messageId
    parameters.Body = body
    parameters.ParseMode = parseMode

    // Make the request
    _, err := http.Post(requestURL, "application/json", strings.NewReader(parameters.AsJSON()))
    if err != nil {
        log.Fatal(err)
    }
}

func (bot Bot) EditMessageKeyboard(chatId int, messageId int, keyboard InlineKeyboardMarkup) {
    requestURL := bot.baseURL + "/editMessageReplyMarkup"
    var parameters EditMessageReplyMarkupRequest
    parameters.ChatId = chatId
    parameters.MessageId = messageId
    parameters.ReplyMarkup = &keyboard

    // Make the request
    _, err := http.Post(requestURL, "application/json", strings.NewReader(parameters.AsJSON()))
    if err != nil {
        log.Fatal(err)
    }
}