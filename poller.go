package telegramBot

import (
    "log"
    "time"
)

func (bot *Bot) Poller() {
    updateOffset := 0

    for {
        <-time.After(time.Second)
        updates := bot.GetUpdates(updateOffset)

        for _, update := range updates {
            bot.HandleUpdate(update)
            updateOffset = update.Id + 1
            log.Println(updateOffset)
        }
    }
}
