package sorter

import (
    "bufio"
    "os"
    "sort"
)

// ExternalSort performs the external sorting of the input file and writes the sorted data to the output file.
func ExternalSort(inputFile, outputFile string) error {
    // Step 1: Split the input file into sorted chunks
    chunkFiles, err := splitIntoChunks(inputFile)
    if err != nil {
        return err
    }

    // Step 2: Merge the sorted chunks into the output file
    err = mergeChunks(chunkFiles, outputFile)
    if err != nil {
        return err
    }

    return nil
}

// splitIntoChunks reads the input file, splits it into manageable chunks, sorts each chunk, and writes them to temporary files.
func splitIntoChunks(inputFile string) ([]string, error) {
    const chunkSize = 10 * 1024 * 1024 // 10 MB
    var chunkFiles []string

    file, err := os.Open(inputFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var strings []string
    for scanner.Scan() {
        str := scanner.Text()
        strings = append(strings, str)

        // If the chunk size exceeds the limit, sort and write to a file
        if len(strings) >= chunkSize/20 { // Assuming average string length is 20 bytes
            chunkFile, err := writeChunk(strings)
            if err != nil {
                return nil, err
            }
            chunkFiles = append(chunkFiles, chunkFile)
            strings = nil // Reset for the next chunk
        }
    }

    // Handle any remaining strings
    if len(strings) > 0 {
        chunkFile, err := writeChunk(strings)
        if err != nil {
            return nil, err
        }
        chunkFiles = append(chunkFiles, chunkFile)
    }

    return chunkFiles, nil
}

// writeChunk sorts the given strings and writes them to a temporary file in the "temp" directory.
func writeChunk(strings []string) (string, error) {
    // Create the temp directory if it doesn't exist
    tempDir := "temp"
    if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
        return "", err
    }

    // Create a temporary file in the temp directory
    chunkFile, err := os.CreateTemp(tempDir, "chunk_*.txt")
    if err != nil {
        return "", err
    }
    defer chunkFile.Close()

    sort.Strings(strings) // Sort the strings
    writer := bufio.NewWriter(chunkFile)
    for _, str := range strings {
        _, err := writer.WriteString(str + "\n")
        if err != nil {
            return "", err
        }
    }
    writer.Flush()

    return chunkFile.Name(), nil
}
