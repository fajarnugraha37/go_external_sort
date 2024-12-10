package main

import (
    "testing"
    "github.com/fajarnugraha37/go_external_sort/sorter"
    "github.com/fajarnugraha37/go_external_sort/utils"
)

func BenchmarkGenerateDataset(b *testing.B) {
	var (
		datasetFile = "input.txt"
		numStrings = 1000000
		stringLength = 16
	)

    for i := 0; i < b.N; i++ {
        err := utils.GenerateDataset(datasetFile, numStrings, stringLength)
        if err != nil {
            b.Fatalf("Error generating dataset: %v", err)
        }
    }
}

func BenchmarkExternalSort(b *testing.B) {
	var (
		datasetFile = "dataset_input.txt"
		outputFile = "dataset_output.txt"
	)

    for i := 0; i < b.N; i++ {
        err := sorter.ExternalSort(datasetFile, outputFile)
        if err != nil {
            b.Fatalf("Sorting failed: %v", err)
        }
    }
}