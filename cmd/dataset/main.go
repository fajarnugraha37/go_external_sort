package main

import (
    "flag"
    "log"
    "github.com/fajarnugraha37/go_external_sort/utils"
)

func main() {
    // Define command-line flags for output file path, number of strings to generate, and string length
    outputFile := flag.String("output", "dataset.txt", "Path to the output file")
    numStrings := flag.Int("count", 100000, "Number of strings to generate")
    stringLength := flag.Int("length", 10, "Length of each random string")
    flag.Parse()

    // Generate the dataset
    err := utils.GenerateDataset(*outputFile, *numStrings, *stringLength)
    if err != nil {
        log.Fatalf("Error generating dataset: %v", err)
    }

    log.Println("Dataset generated successfully.")
}