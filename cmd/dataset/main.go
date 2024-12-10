package main

import (
    "flag"
    "log"
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

func main() {
    // Define command-line flags for output file path, number of strings to generate, and string length
    outputFile := flag.String("output", "dataset.txt", "Path to the output file")
    numStrings := flag.Int("count", 100000, "Number of strings to generate")
    stringLength := flag.Int("length", 10, "Length of each random string")
    flag.Parse()

    // Generate the dataset
    err := generateDataset(*outputFile, *numStrings, *stringLength)
    if err != nil {
        log.Fatalf("Error generating dataset: %v", err)
    }

    log.Println("Dataset generated successfully.")
}

// generateDataset creates a file with the specified number of random strings.
func generateDataset(outputFile string, count int, length int) error {
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
