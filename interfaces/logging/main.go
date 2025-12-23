package main

import (
	"fmt"
	"net/http"
	"os"
)

type logger interface {
	Log(message string)
}
type infologger struct {
}

type errorlogger struct {
}

type filelogger struct {
	filename string
}

func main() {
	info := infologger{}
	errorlog := errorlogger{}
	filelogger_v1 := filelogger{filename: "log.txt"}

	writeLog(info, "server started")
	writeLog(errorlog, "could not connect to DB")
	writeLog(filelogger_v1, "server started")
	writeLog(filelogger_v1, "user logged in")
	fmt.Println("Logs written to log.txt")

	resp, err := http.Get("httpL//google.com")
	if err != nil {
		fmt.Println("error:", err)
	}
	bs := make([]byte, 1000)
	resp.Body.Read(bs)
}
func (i infologger) Log(message string) {
	fmt.Println("INFO:", message)
}

func (e errorlogger) Log(message string) {
	fmt.Println("ERROR:", message)
}

func (f filelogger) Log(message string) {

	// now file becomes a variable of type pointer to file *File
	file, err := os.OpenFile(f.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	// WriteString is method, which needs to be called with an variable of type pointerToFile
	// Above, file variable is of type pointerToFile, so we are calling WriteString with 'file' variabel
	file.WriteString("file log: " + message + "\n")
}

func writeLog(l logger, msg string) {
	l.Log(msg)
}
