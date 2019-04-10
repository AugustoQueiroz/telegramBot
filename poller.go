package telegramBot

import (
    "log"
    "time"
)

func Poller() {
    updateOffset := 0

    for {
        <-time.After(time.Second)
        updates := GetUpdates(updateOffset)

        for _, update := range updates {
            HandleUpdate(update)
            updateOffset = update.Id + 1
            log.Println(updateOffset)
        }
    }
}
