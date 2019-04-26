package telegramBot

import (
    // Standard Packages
    _ "os"
    _ "log"
    _ "strings"
    _ "strconv"
    _ "net/url"
    _ "net/http"
    _ "encoding/json"
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

// Assigns a function to a given command
// - Parameter command: The command to be recognized exactly as it is meant to be read
// - Parameter function: The function to be executed when the command is recognized in a message
func (bot Bot) HandleFunc(command string, function func(*Message)) {
    bot.commandHandlers[command] = function
    bot.commandHandlers[command + "@SecretSantainatorBot"] = function // TODO: Change this to take the bot's username
}
