package tools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GetPath(localPath string) string {
	pathToRootOfProgram, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}
	path := filepath.FromSlash(fmt.Sprintf("%s/%s", pathToRootOfProgram, localPath))
	return path
}

func ReadTextFile(path string) string {
	file, err := os.Open(path)
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err.Error())
		}
	}()
	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(file)

	var text string
	for scanner.Scan() {
		text += fmt.Sprintf("%v\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	return text
}

func WriteTextFile(path, text string) {
	data := []byte(text)
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func CheckAndCreateFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		WriteTextFile(path, "")
	}
}

func SaveToLog(err error, debugStack []byte) {
	path := GetPath("log.txt")
	CheckAndCreateFile(path)

	currentTime := getCurrentDateAndTime()
	oldLogText := ReadTextFile(path)
	text := fmt.Sprintf("[%s]: %s\n%s\n%s", currentTime, err.Error(), string(debugStack), oldLogText)

	WriteTextFile(path, text)
}

func getCurrentDateAndTime() string {
	return convertUnixTimeToDate(int(time.Now().Unix()))
}

func convertUnixTimeToDate(ut int) string {
	t := time.Unix(int64(ut), 0)
	dateFormat := "02.01.2006 15:04:05"
	date := t.Format(dateFormat)
	return date
}

func ReceiveUserInput() string {
	var userInput string
	in := bufio.NewReader(os.Stdin)
	userInput, err := in.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	if len(userInput) > 0 {
		u := strings.Split(userInput, "")
		u = u[:len(u)-1]
		userInput = strings.Join(u, "")
	}

	return userInput
}
