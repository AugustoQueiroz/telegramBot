package telegramBot

import (
    // Standard Packages
    _ "os"
    "log"
    _ "strings"
    _ "strconv"
    _ "net/url"
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

    id              int
    username        string

    // Public Attributes
    CallbackHandler callbackHandler
}

func NewBot(token string) (bot Bot) {
    bot.token = token
    bot.baseURL = "https://api.telegram.org/bot" + token
    bot.commandHandlers = make(map[string]commandHandler)

    bot.GetMe()
    log.Println(bot.username)

    return
}

// Get the information about this bot
func (bot *Bot) GetMe() {
    requestURL := bot.baseURL + "/getMe"

    response, err := http.Get(requestURL)
    if err != nil {
        log.Println("Error getting the information about the bot", err)
    }

    var responseBody map[string]interface{}
    json.NewDecoder(response.Body).Decode(&responseBody)

    botInfo := responseBody["result"].(map[string]interface{})

    bot.id = int(botInfo["id"].(float64))
    bot.username = botInfo["username"].(string)
}

// Check wether a given token is this bot's token
// - Parameter inputToken: The token being to be tested against the bot token
// - Returns: True if the received token is the bots token
func (bot *Bot) CheckToken(inputToken string) bool {
    return inputToken == bot.token
}

// Assigns a function to a given command
// - Parameter command: The command to be recognized exactly as it is meant to be read
// - Parameter function: The function to be executed when the command is recognized in a message
func (bot *Bot) HandleFunc(command string, function func(*Message)) {
    bot.commandHandlers[command] = function
    bot.commandHandlers[command + "@SecretSantainatorBot"] = function // TODO: Change this to take the bot's username
}
