package sorter

import (
    "bufio"
    "os"
    "sort"
)

// chunk represents a sorted chunk with its associated file and scanner.
type chunk struct {
    file    *os.File
    scanner *bufio.Scanner
    value   string
}

// priorityQueue is a simple min-heap for managing chunks.
type priorityQueue []*chunk

// newPriorityQueue creates a new priority queue.
func newPriorityQueue() *priorityQueue {
    return &priorityQueue{}
}

// push adds a chunk to the priority queue.
func (pq *priorityQueue) push(c *chunk) {
    *pq = append(*pq, c)
    sort.Slice(*pq, func(i, j int) bool {
        return (*pq)[i].value < (*pq)[j].value
    })
}

// pop removes and returns the smallest chunk from the priority queue.
func (pq *priorityQueue) pop() *chunk {
    min := (*pq)[0]
    *pq = (*pq)[1:]
    return min
}

// len returns the number of chunks in the priority queue.
func (pq *priorityQueue) len() int {
    return len(*pq)
}


// mergeChunks merges the sorted chunk files into a single output file.
func mergeChunks(chunkFiles []string, outputFile string) error {
    // Create the output file
    outFile, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer outFile.Close()

    // Create a priority queue to manage the sorted chunks
    pq := newPriorityQueue()

    // Open all chunk files and initialize the priority queue
    for _, chunkFile := range chunkFiles {
        file, err := os.Open(chunkFile)
        if err != nil {
            return err
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        if scanner.Scan() {
            str := scanner.Text()
            pq.push(&chunk{file: file, scanner: scanner, value: str})
        }
    }

    // Merge the chunks
    writer := bufio.NewWriter(outFile)
    defer writer.Flush()

    for pq.len() > 0 {
        // Get the smallest element from the priority queue
        minChunk := pq.pop()
        _, err := writer.WriteString(minChunk.value + "\n")
        if err != nil {
            return err
        }

        // Read the next value from the chunk
        if minChunk.scanner.Scan() {
            str := minChunk.scanner.Text()
            minChunk.value = str
            pq.push(minChunk)
        } else {
            // Close the file if there are no more values
            minChunk.file.Close()
        }
    }

    return nil
}