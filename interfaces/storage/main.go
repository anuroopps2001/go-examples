package main

import (
	"fmt"
	"os"
)

type Storage interface {
	// Any type with a method Save(string) error is Storage.
	Save(data string) error
}

type consoleStorage struct{}

type memoryStorage struct {
	items []string
}

type fileStorage struct {
	filename string
}

func main() {
	consle := consoleStorage{}
	fstorage := fileStorage{filename: "os.txt"}
	memoryStorage := &memoryStorage{}
	persist(consle, "Hello from console")
	persist(fstorage, "Hello from file storage")
	persist(memoryStorage, "hello from memory storage")
	fmt.Println(memoryStorage.items)

}

func (c consoleStorage) Save(data string) error {
	fmt.Println("Console:", data)
	return nil
}
func (m *memoryStorage) Save(data string) error {
	m.items = append(m.items, data)
	return nil
}
func (f fileStorage) Save(data string) error {
	file, err := os.OpenFile(f.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(data + "\n")
	return err

}

func persist(s Storage, data string) {
	err := s.Save(data)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
