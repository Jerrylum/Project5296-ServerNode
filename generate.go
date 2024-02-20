package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func writeFileWithAs(filename string, numOfAs int) error {
	// Create the file path
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(currentDir, "public", "download", filename)

	fd, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer fd.Close()

	fd.WriteAt([]byte("a"), int64(numOfAs-1))

	return nil
}

func main() {
	// Define and parse the command-line argument
	sizeMB := flag.Float64("m", 0, "Generated Size in MB")
	flag.Parse()

	if *sizeMB <= 0 {
		fmt.Println("Please specify the size in MB")
		return
	}

	// Call the function to write the file
	finalFileName := float64(int(*sizeMB*1000000)) / float64(1000000)
	finalSizeMB := int(*sizeMB * 1000 * 1000)
	err := writeFileWithAs(fmt.Sprintf("%s.out", strconv.FormatFloat(finalFileName, 'f', -1, 64)), finalSizeMB)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}
