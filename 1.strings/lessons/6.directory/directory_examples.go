package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CheckLogFile() {
	if _, err := os.Stat("log.txt"); err == nil {
		fmt.Println("Log.txt file exists")
	}
}

func CheckLogNotFile() {
	// Check if log.txt file exists
	if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
		fmt.Println("Log.txt file does not exist")
	} else {
		fmt.Println("Log.txt file exists")
	}

	// Check if log1.txt file exists
	if _, err := os.Stat("log1.txt"); os.IsNotExist(err) {
		fmt.Println("Log1.txt file does not exist")
	} else {
		fmt.Println("Log1.txt file exists")
	}
}

func ReadFile() {
	contentBytes, err := os.ReadFile("/data/names.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var contentStr string = string(contentBytes)
	fmt.Println(contentStr)
}

func WriteFile() {
	content := "Hello, World!\nThis is a test file.\n"
	err := os.WriteFile("/data/hello_world.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")
}

func WriteTemporaryFile() {
	tempFile, err := os.CreateTemp("", "tempfile_*.txt")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up the temporary file after use

	content := "This is a temporary file.\n"
	if _, err := tempFile.Write([]byte(content)); err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}

	fmt.Println("Temporary file created:", tempFile.Name())
}

func CountLinesInFile() {
	file, _ := os.Open("/data/names.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	defer file.Close()
	fmt.Println(lineCount)
}

func ReadLine(lineNumber int) string {
	file, _ := os.Open("/data/names.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	defer file.Close()
	return ""
}

func CompareFiles() bool {
	file1Content, err1 := os.ReadFile("/data/one.txt")
	if err1 != nil {
		fmt.Println("Error reading file:", err1)
		return false
	}

	file2Content, err2 := os.ReadFile("/data/two.txt")
	if err2 != nil {
		fmt.Println("Error reading file:", err2)
		return false
	}

	return string(file1Content) == string(file2Content)
}

func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully")
}

func CopyFile() {
	original, err := os.Open("data/original.txt")
	if err != nil {
		panic(err)
	}
	defer original.Close()
	original_copy, err2 := os.Create("data/copy.txt")
	if err2 != nil {
		panic(err2)
	}
	defer original_copy.Close()
	_, err3 := io.Copy(original_copy, original)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("File copied successfully")
}

func RenameFile() {
	err := os.Rename("data/copy.txt", "data/moved.txt")
	if err != nil {
		fmt.Println("Error moving file:", err)
		return
	}
	fmt.Println("File renamed successfully")
}

func MoveFile() {
	// create "target" directory if it doesn't exist
	if _, err := os.Stat("target"); os.IsNotExist(err) {
		err := os.Mkdir("target", 0755)
		if err != nil {
			fmt.Println("Error creating target directory:", err)
			return
		}
	}
	// move the file to the "target" directory
	original, err := os.Open("data/moved.txt")
	if err != nil {
		panic(err)
	}
	defer original.Close()
	target, err := os.Create("target/moved.txt")
	if err != nil {
		panic(err)
	}
	defer target.Close()
	_, err = io.Copy(target, original)
	if err != nil {
		panic(err)
	}
	// delete the original file after moving
	err = os.Remove("data/moved.txt")
	if err != nil {
		fmt.Println("Error deleting original file:", err)
		return
	}
	fmt.Println("File moved to target directory successfully")

}

func ListFilesInDirectory() {
	files, err := os.ReadDir("data")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
