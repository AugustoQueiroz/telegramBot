package telegramBot

import (
    // Standard Packages
    _ "os"
    "log"
    _ "strings"
    "strconv"
    "net/url"
    "net/http"
    "encoding/json"
)

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
    var responseBody UpdatesResponse
    json.NewDecoder(response.Body).Decode(&responseBody)

    updates = responseBody.Updates

    return
}