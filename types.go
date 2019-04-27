package telegramBot

import (
    "io"
    "encoding/json"
)

type UpdatesResponse struct {
    Ok                  bool                `json:"ok"`
    Updates             []Update            `json:"result"`
}

type MessageResponse struct {
    Ok                  bool                `json:"ok"`
    Result              *Message            `json:"result"`
}

type Update struct {
    // Mandatory Attributes
    Id                  int                 `json:"update_id"`                  // Unique identifier for the update

    // Optional Attributes
    Message             *Message            `json:"message,omitempty"`          // New incoming message
    CallbackQuery       *CallbackQuery      `json:"callback_query,omitempty"`
}

type CallbackQuery struct {
    // Mandatory Attributes
    Id                  int                 `json:"id"`                         //
    From                *User               `json:"from"`                       //
    ChatInstance        string              `json:"chat_instance"`              //

    // Optional Attributes
    Message             *Message            `json:"message,omitempty"`          //
    InlineMessageId     string              `json:"inline_message_id,omitempty"`
    Data                string              `json:"data,omitempty"`
}

type User struct {
    // Mandatory Attributes
    Id                  int                 `json:"id"`                         // Unique identifier for the user
    IsBot               bool                `json:"is_bot"`                     // Wether or not the user is a bot
    FirstName           string              `json:"first_name"`                 // First name of the user/bot

    // Optional Attributes
    LastName            string              `json:"last_name,omitempty"`        // Last name of the user/bot
    Username            string              `json:"username,omitempty"`         // Username of the user/bot
    LanguageCode        string              `json:"language_code,omitempty"`    // Code of the user's language
}

type Chat struct {
    // Mandatory Attributes
    Id                  int                 `json:"id"`                                             // Unique identifier of the chat
    Type                string              `json:"type"`                                           // "private", "group", "supergroup", "channel"

    // Optional Attributes
    Title               string              `json:"title,omitempty"`                                // The title of the group/supergroup/channel
    Username            string              `json:"username,omitempty"`                             // User name of private/supergroup/channel
    FirstName           string              `json:"first_name,omitempty"`                           // First name of the other party on private chat
    LastName            string              `json:"last_name,omitempty"`                            // Last name of the other party on private chat
    AllMembersAreAdm    string              `json:"all_members_are_administrators,omitempty"`       // Whether or not all the members of the group are adms
}

type Message struct {
    // Mandatory Attributes
    Id                  int                 `json:"message_id"`                                     // Unique message identifier
    Date                int                 `json:"date"`                                           // Date the message was sent Unix time
    Origin              *Chat               `json:"chat"`                                           // Chat the message belongs to

    // Optional Attributes
    From                *User               `json:"from,omitempty"`                                 // Sender, empty when sent to channels
    Body                string              `json:"text,omitempty"`                                  // Body of the message
    Entities            []MessageEntity     `json:"entities,omitempty"`                             // Entities such as usernames, URLs, bot commands, etc
}

type MessageEntity struct {
    Type                string              `json:"type"`                                           // "hashtag", "cashtag", "bot_command", "url", "email", "phone_number", "bold", "italic", "code", "pre", "text_link", or "text_mention"
    Offset              int                 `json:"offset"`                                         // Offset to the start of the entity in UTF-16 code units
    Length              int                 `json:"length"`                                         // Length of the entity in UTF-16 code units
    URL                 string              `json:"url,omitempty"`                                  // URL if "text_link"
    User                *User               `json:"user,omitempty"`                                 // Mentioned user if "text_mention"
}

type InlineKeyboardMarkup struct {
    Keyboard            [][]InlineKeyboardButton  `json:"inline_keyboard"`                            // Matrix of the buttons to be presented
}

type InlineKeyboardButton struct {
    // Mandatory Attributes
    Label               string              `json:"text"`                                           // Label on the button

    // Optional Attributes
    URL                 string              `json:"url,omitempty"`                                  // HTTP or tg:// url to be opened by button press
    CallbackData        string              `json:"callback_data,omitempty"`                        // Data to be sent in callback to bot on button press
}

func DecodeUpdate(body io.ReadCloser) (update Update) {
    update = Update{}

    decoder := json.NewDecoder(body)
    err := decoder.Decode(&update)
    if err != nil {
        panic(err)
    }

    return
}
