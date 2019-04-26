package telegramBot

import (
    // Standard Packages
    _ "os"
    "log"
    "strings"
    "strconv"
    "net/url"
    "net/http"
    "encoding/json"
)

type commandHandler func(*Message)
type callbackHandler func(*CallbackQuery)

type Bot struct {
    // Private Attributes
    token           string
    baseURL          string
    commandHandlers map[string]commandHandler

    username        string

    // Public Attributes
    CallbackHandler callbackHandler
}

func NewBot(token string) (bot Bot) {
    bot.token = token
    bot.baseURL = "https://api.telegram.org/bot" + token
    bot.commandHandlers = make(map[string]commandHandler)

    return
}

// Check wether a given token is this bot's token
// - Parameter inputToken: The token being to be tested against the bot token
// - Returns: True if the received token is the bots token
func (bot Bot) CheckToken(inputToken string) bool {
    return inputToken == bot.token
}

// Sets the bot webhook and returns whether or not it was successful
// - Parameter webhookURL: The url that updates should be sent to
// - Returns: Whether or not the creation of the webhook was successful
func (bot Bot) SetWebhook(webhookURL string) bool {
    // Create the request
    requestURL := bot.baseURL + "/setWebhook"
    parameters := url.Values {
        "url": {webhookURL + bot.token + "/"},
    }

    // Make the request
    response, err := http.PostForm(requestURL, parameters)
    if err != nil {
        log.Fatal(err)
    }

    // Parse the response
    var responseBody map[string]interface{}
    json.NewDecoder(response.Body).Decode(&responseBody)

    return responseBody["result"].(bool)
}

// Deletes the webhook and returns whether or not was successfull
// - Returns: True if the deletion of the webhook was successful
func (bot Bot) DeleteWebhook() bool {
    // Create the request
    requestURL := bot.baseURL + "/deleteWebhook"

    // Make the request
    response, err := http.Get(requestURL)
    if err != nil {
        log.Fatal(err)
    }

    // Parse the response
    var responseBody map[string]interface{}
    json.NewDecoder(response.Body).Decode(&responseBody)

    return responseBody["result"].(bool)
}

// Gets the updates manually
// - Parameter offset: The offset in the updates to be received
// - Returns: A slice of Update objects
func (bot Bot) GetUpdates(offset int) (updates []Update) {
    // Create the request
    requestURL := bot.baseURL + "/getUpdates"
    parameters := url.Values {
        "offset": {strconv.Itoa(offset)},
        "timeout": {"1"},
        //"allowed_updates": {"message", "callback_query"},
    }

    // Make the request
    response, err := http.PostForm(requestURL, parameters)
    if err != nil {
        log.Fatal(err)
    }

    // Parse the response
    var responseBody Response
    json.NewDecoder(response.Body).Decode(&responseBody)

    updates = responseBody.Result

    return
}

// Receives an update and checks whether or not it has one of the known commands
// Then calls the function for that command
// - Parameter update: The update that will be checked for commands
func (bot Bot) HandleUpdate(update Update) {
    if update.Message == nil {
        // Handle Callback Query
        callback := update.CallbackQuery
        bot.CallbackHandler(callback)
        return
    }
    message := update.Message

    if len(message.Entities) > 0 {
        // If the message received has at least one entity (that can be commands, usernames, etc)
        // then will check to see if any of them is a recognized command
        for _, entity := range message.Entities {
            if entity.Type == "bot_command" {
                command := ExtractEntity(message.Body, entity.Offset, entity.Length)

                handler, isDefined := bot.commandHandlers[command]
                if isDefined {
                    // If the command is recognized (aka, has been assigned a handler function)
                    handler(message) // Calls that function
                }
            }
        }
    }
}

// Assigns a function to a given command
// - Parameter command: The command to be recognized exactly as it is meant to be read
// - Parameter function: The function to be executed when the command is recognized in a message
func (bot Bot) HandleFunc(command string, function func(*Message)) {
    bot.commandHandlers[command] = function
    bot.commandHandlers[command + "@SecretSantainatorBot"] = function // TODO: Change this to take the bot's username
}

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
