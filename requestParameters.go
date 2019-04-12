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

type EditMessageTextRequest struct {
    ChatId                  int                             `json:"chat_id,omitempty"`
    MessageId               int                             `json:"message_id,omitempty"`
    Body                    string                          `json:"text"`
    ParseMode               string                          `json:"parse_mode,omitempty"`
//    ReplyMarkup             *InlineKeyboardMarkup            `json:"reply_markup,omitempt"`
}

type EditMessageReplyMarkupRequest struct {
    ChatId                  int                             `json:"chat_id"`
    MessageId               int                             `json:"message_id"`
    ReplyMarkup             *InlineKeyboardMarkup           `json:"reply_markup"`
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

func (request EditMessageTextRequest) AsJSON() (JSONString string) {
    JSONObject, err := json.Marshal(&request)
    if err != nil {
        log.Fatal("Error trying to jsonify EditMessageTextRequest", err)
    }

    JSONString = string(JSONObject)
    return
}



func (request EditMessageReplyMarkupRequest) AsJSON() (JSONString string) {
    JSONObject, err := json.Marshal(&request)
    if err != nil {
        log.Fatal("Error trying to jsonify EditMessageReplyMarkupRequest", err)
    }

    JSONString = string(JSONObject)
    return
}
