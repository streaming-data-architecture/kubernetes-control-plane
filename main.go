package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
	_ "strconv"
	"time"
)

func main() {
	for {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		filename := "output.md"
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		content := []byte(currentTime + "\n")
		if _, err := file.Write(content); err != nil {
			fmt.Println(err)
		}
		fmt.Println(currentTime)
		time.Sleep(5 * time.Second)
	}
}
