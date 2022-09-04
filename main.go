package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const plags = os.O_CREATE | os.O_WRONLY

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(workingDir, "README.md")

	newPath, err := copyFileToTempDir(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Your new file path is: %s\n", newPath)

	err = printFileToTerminal(newPath)
	if err != nil {
		panic(err)
	}

}

// copyFileToTempDir copies the `filePath` into the temp directory and keep the original filename.
func copyFileToTempDir(filePath string) (string, error) {
	fileName := filepath.Base(filePath)
	if fileName == "." {
		// filePath is empty
		return "", fmt.Errorf("filePath is empty")
	}
	newPath := filepath.Join(os.TempDir(), fileName)

	// Open the source file
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer sourceFile.Close()

	// Open the target file
	targetFile, err := os.OpenFile(newPath, plags, 0644)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	// Stream the sourceFile's content directly into the targetFile
	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		return "", err
	}

	return newPath, nil
}

// printFileToTerminal prints the filePath's content to the terminal.
func printFileToTerminal(filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Stream the file's content to stdout to show it worked
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		return err
	}
	return nil
}
