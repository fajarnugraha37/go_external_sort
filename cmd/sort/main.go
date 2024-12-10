package main

import (
    "flag"
    "log"
    "github.com/fajarnugraha37/go_external_sort/sorter"
)

func main() {
    // Define command-line flags for input and output file paths
    inputFile := flag.String("input", "", "Path to the input file")
    outputFile := flag.String("output", "", "Path to the output file")
    flag.Parse()

    // Validate input and output file paths
    if *inputFile == "" || *outputFile == "" {
        log.Fatal("Input and output file paths must be specified.")
    }

    // Perform external sorting
    err := sorter.ExternalSort(*inputFile, *outputFile)
    if err != nil {
        log.Fatalf("Error during sorting: %v", err)
    }

    log.Println("Sorting completed successfully.")
}
