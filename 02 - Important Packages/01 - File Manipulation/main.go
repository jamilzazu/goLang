package main

import (
	"bufio"
	"fmt"
	"os"
)

func createNewFile() {
	var file = createFile()
	writeStringToFile(file)
	writeByteToFile(file)
	closeFile(file)
	readFile(file.Name())
	readFileInParts(file.Name())
	removeFile(file.Name())
}

func createFile() *os.File {
	file, err := os.Create("file.txt")
	messageError(err)
	return file
}

func removeFile(fileName string) {
	err := os.Remove(fileName)
	messageError(err)
}

func closeFile(file *os.File) {
	file.Close()
}

func readFile(fileName string) {
	file, err := os.ReadFile(fileName)
	messageError(err)
	fmt.Println("\nReading file:")
	fmt.Printf(string(file))
}

func readFileInParts(fileName string) {
	file, err := os.Open(fileName)
	messageError(err)
	fmt.Println("\nused when the file size is too large")
	fmt.Println("Reading file in parts, 10 byte size:")
	readBufferAndPrint(file)
	closeFile(file)
}

func readBufferAndPrint(file *os.File) {
	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)
	for {
		position, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:position]))
	}
}

func writeStringToFile(file *os.File) {
	fmt.Println("For strings, use file.WriteString")
	size, err := file.WriteString("Hello World String! \n")
	messageError(err)
	fmt.Printf("File created successfully, size: %d bytes \n", size)
}

func writeByteToFile(file *os.File) {
	fmt.Println("\nFor any type, file.Write")
	size, err := file.Write([]byte("Hello World Bytes 1234455! \n"))
	messageError(err)
	fmt.Printf("File created successfully, size: %d bytes \n", size)
}

func messageError(message error) {
	if message != nil {
		panic(message)
	}
}

func main() {
	createNewFile()
}
