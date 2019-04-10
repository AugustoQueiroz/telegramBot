package telegramBot

import (
    // Standard Packages
    "os"
    "log"
    "strings"
    "strconv"
    "net/url"
    "net/http"
    "encoding/json"

    // External Packages
    "github.com/gorilla/mux"
)

type commandHandler func(*Message)

var (
    // Private Attributes
    token           string
    apiURL          string
    commandHandlers map[string]commandHandler

    // Public Attributes
    CallbackHandler func(*CallbackQuery)
)

func init() {
    token = os.Getenv("TELEGRAM_TOKEN")
    if token == "" {
        log.Fatal("$TELEGRAM_TOKEN was not set")
    }

    apiURL = "https://api.telegram.org/bot" + token

    commandHandlers = make(map[string]commandHandler)
}

// Check wether a given token is this bot's token
func CheckToken(inputToken string) bool {
    return inputToken == token
}

// Sets the bot webhook and returns whether or not it was successful
func SetWebhook(webhookURL string) bool {
    // Create the request
    requestURL := apiURL + "/setWebhook"
    parameters := url.Values {
        "url": {webhookURL + token + "/"},
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
func DeleteWebhook() bool {
    // Create the request
    requestURL := apiURL + "/deleteWebhook"

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
func GetUpdates(offset int) (updates []Update) {
    // Create the request
    requestURL := apiURL + "/getUpdates"
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
func HandleUpdate(update Update) {
    if update.Message == nil {
        // Handle Callback Query
        callback := update.CallbackQuery
        CallbackHandler(callback)
        return
    }
    message := update.Message

    if len(message.Entities) > 0 {
        // If the message received has at least one entity (that can be commands, usernames, etc)
        // then will check to see if any of them is a recognized command
        for _, entity := range message.Entities {
            if entity.Type == "bot_command" {
                command := ExtractEntity(message.Body, entity.Offset, entity.Length)

                handler, isDefined := commandHandlers[command]
                if isDefined {
                    // If the command is recognized (aka, has been assigned a handler function)
                    handler(message) // Calls that function
                }
            }
        }
    }
}

// Handles the updates received by webhooks
func HandleUpdates(writer http.ResponseWriter, request* http.Request) {
    if mux.Vars(request)["token"] == token {
        update := DecodeUpdate(request.Body)

        HandleUpdate(update)
    }
}

// Assigns a function to a given command
func HandleFunc(command string, function func(*Message)) {
    commandHandlers[command] = function
    commandHandlers[command + "@SecretSantainatorBot"] = function // TODO: Change this to take the bot's username
}

// Sends a message with the given parameters
func SendMessageWithParameters(message MessageRequest) Message {
    // Create the request
    requestURL := apiURL + "/sendMessage"


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
func SendMessage(body string, chatId int, parseMode string) Message {
    // Define the parameters
    var message MessageRequest
    message.ChatId = chatId
    message.Body = body
    message.ParseMode = parseMode

    return SendMessageWithParameters(message)
}

// Sends a text message to a given chat with an inline keyboard
func SendMessageWithKeyboard(body string, chatId int, parseMode string, replyMarkup InlineKeyboardMarkup) Message {
    // Define the parameters
    var message MessageRequest
    message.ChatId = chatId
    message.Body = body
    message.ParseMode = parseMode
    message.ReplyMarkup = &replyMarkup

    return SendMessageWithParameters(message)
}

// Sends a Markdown message to the given chat
func SendMarkdownMessage(body string, chatId int) Message {
    return SendMessage(body, chatId, "Markdown")
}

// Sends an HTML message to the given chat
func SendHTMLMessage(body string, chatId int) Message {
    return SendMessage(body, chatId, "HTML")
}

func EditMessageText(chatId int, messageId int, body string, parseMode string) {
    // Create the request
    requestURL := apiURL + "/editMessageText"
    var parameters EditMessageRequest
    parameters.ChatId = chatId
    parameters.MessageId = messageId
    parameters.Body = body
    parameters.ParseMode = parseMode

    // Make the request
    log.Println(parameters.AsJSON())
    response, err := http.Post(requestURL, "application/json", strings.NewReader(parameters.AsJSON()))
    if err != nil {
        log.Fatal(err)
    }

    log.Println(response)
}
