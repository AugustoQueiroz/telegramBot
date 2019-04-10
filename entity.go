package telegramBot

func ExtractEntity(message string, position int, length int) (entity string) {
    // UTF-16 uses 16 bits code units, so a code unit is 2 bytes
    // So if the position is x in UTF-16 code units, it will be 2x in bytes
    // And if the end is y UTF-16 code units , in bytes it will be 2y
    message_runes := []rune(message)
    start := position
    end := (position + length)
    entity = string(message_runes[start:end])
    return
}
