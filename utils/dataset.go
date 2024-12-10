package utils

import (
    "math/rand"
    "os"
    "time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generates a random string of a specified length.
func RandomString(length int) string {
    b := make([]byte, length)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

// generateDataset creates a file with the specified number of random strings.
func GenerateDataset(outputFile string, count int, length int) error {
    // Create the output file
    file, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer file.Close()

    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Write random strings to the file
    for i := 0; i < count; i++ {
        _, err := file.WriteString(RandomString(length) + "\n")
        if err != nil {
            return err
        }
    }

    return nil
}
