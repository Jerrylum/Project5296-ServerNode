package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func writeFileWithAs(filename string, numOfAs int) error {
	// Create the file path
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(currentDir, "public", "download", filename)

	// Create the file content
	// content := make([]byte, numOfAs)
	// for i := range content {
	// 	content[i] = 'a'
	// }

	// Write the file
	// return os.WriteFile(filePath, content, 0644)

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
	sizeMB := flag.Int("m", 0, "Generated Size in MB")
	flag.Parse()

	if *sizeMB == 0 {
		fmt.Println("Please specify the size in MB")
		return
	}

	// Call the function to write the file
	err := writeFileWithAs(fmt.Sprintf("%d.out", *sizeMB), *sizeMB*1000*1000)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}
