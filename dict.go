package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func readDictionary() []byte {
	//open dictionary file
	file, err := os.Open("dictionary")
	checkError(err)
	defer file.Close()

	//get the file size
	stat, err := file.Stat()
	checkError(err)

	//read the file
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	checkError(err)

	return data
}

func searchWordLine(word string) string {
	data := string(readDictionary())
	wordLines := strings.Split(data, "\n")
	var wordLine string
	for i := 0; i < len(wordLines); i++ {
		target := strings.Split(wordLines[i], ":")[0]
		if strings.ToLower(word) == strings.ToLower(target) {
			wordLine = wordLines[i]
		}
	}
	return wordLine
}

func printWordLine(wordLine string) {
	for i := 0; i < len(strings.Split(wordLine, ":")); i++ {
		fmt.Println(strings.Split(wordLine, ":")[i])
	}
	fmt.Println("_______________________________")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Open Source Dictionary\n")
	for {
		fmt.Print("(OSD1.1):>")
		line, _ := reader.ReadString('\n')
		word := strings.Split(line, "\n")[0]
		wordLine := searchWordLine(word)
        if wordLine != "" {
            printWordLine(wordLine)
        } else {
            fmt.Println("Oooooooo!")
        }
	}
}
