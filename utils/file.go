package utils

import (
    "os"
)

// CheckError checks for an error and logs it if present.
func CheckError(err error) {
    if err != nil {
        // Log the error (you can replace this with a proper logging mechanism)
        panic(err)
    }
}

// RemoveTempFiles removes temporary files created during sorting.
func RemoveTempFiles(files []string) {
    for _, file := range files {
        os.Remove(file)
    }
}
