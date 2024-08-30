package main

import (
	"fmt"
	"os"
	"patterns/pkg/yalm"
)

func main() {
	key := os.Getenv("API_YALM")
	url := os.Getenv("URL_YALM")
	folderId := os.Getenv("FOLDERID_YALM")

	gpt := yalm.NewYaLM(key, url, folderId)

	answer, err := gpt.Promt("Переведи текст", "Hi, world!")
	if err != nil {
		panic(err)
	}

	fmt.Println(answer)
}
