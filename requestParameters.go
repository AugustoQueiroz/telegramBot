package telegramBot

import (
    "log"
    "encoding/json"
)

type MessageRequest struct {
    ChatId                  int                             `json:"chat_id"`
    Body                    string                          `json:"text"`
    ParseMode               string                          `json:"parse_mode,omitempty"`
    DisableWebPagePreview   bool                            `json:"disable_web_page_preview,omitempty"`
    DisableNotification     bool                            `json:"disable_notification,omitempty"`
    ReplyToMessageId        int                             `json:"reply_to_message_id,omitempty"`
    ReplyMarkup             *InlineKeyboardMarkup            `json:"reply_markup,omitempty"`
}

type EditMessageRequest struct {
    ChatId                  int                             `json:"chat_id,omitempty"`
    MessageId               int                             `json:"message_id,omitempty"`
    Body                    string                          `json:"text"`
    ParseMode               string                          `json:"parse_mode,omitempty"`
//    ReplyMarkup             *InlineKeyboardMarkup            `json:"reply_markup,omitempt"`
}

// Return the json string of a MessageRequest
func (request MessageRequest) AsJSON() (JSONString string) {
    JSONObject, err := json.Marshal(&request)
    if err != nil {
        log.Fatal("Error trying to jsonify MessageRequest", err)
    }

    JSONString = string(JSONObject)
    return
}

func (request EditMessageRequest) AsJSON() (JSONString string) {
    JSONObject, err := json.Marshal(&request)
    if err != nil {
        log.Fatal("Error trying to jsonify EditMessageRequest", err)
    }

    JSONString = string(JSONObject)
    return
}
