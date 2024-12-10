package utils

import (
    "log"
)

// LogInfo logs informational messages.
func LogInfo(message string) {
    log.Println("INFO:", message)
}

// LogError logs error messages.
func LogError(message string) {
    log.Println("ERROR:", message)
}
