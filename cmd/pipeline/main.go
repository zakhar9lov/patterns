package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"patterns/pkg/yalm"
	"strings"
)

type Job struct {
	LoadPath string
	Text     string
	SavePath string
}

func main() {

	key := os.Getenv("API_YALM")
	url := os.Getenv("URL_YALM")
	folderId := os.Getenv("FOLDERID_YALM")

	gpt := yalm.NewYaLM(key, url, folderId)

	paths, err := filepath.Glob("input/*.txt")
	if err != nil {
		log.Fatalf("File search error: %v", err)
	}

	ch1 := loadFiles(paths)
	ch2 := translateFiles(gpt, ch1)
	ch3 := toUpperCase(ch2)
	result := saveFiles(ch3)

	for ok := range result {
		if ok {
			fmt.Println("DONE!")
		} else {
			fmt.Println("Failed!")
		}
	}

}

func loadFiles(paths []string) <-chan Job {
	ch := make(chan Job)

	go func() {
		defer close(ch)

		for _, v := range paths {
			text, err := os.ReadFile(v)
			if err != nil {
				return
			}

			job := Job{
				LoadPath: v,
				Text:     string(text),
				SavePath: strings.Replace(v, "input", "output", 1),
			}

			ch <- job
		}
	}()

	return ch
}

func translateFiles(gpt *yalm.YaLM, chIn <-chan Job) <-chan Job {
	chOut := make(chan Job)

	go func() {
		defer close(chOut)

		for job := range chIn {
			res, err := gpt.Promt("Переведи текст", job.Text)
			if err != nil {
				log.Printf("Fetching error: %v", err)
				return
			}
			job.Text = res
			chOut <- job
		}

	}()

	return chOut
}

func toUpperCase(chIn <-chan Job) <-chan Job {
	chOut := make(chan Job)

	go func() {
		defer close(chOut)

		for job := range chIn {
			job.Text = strings.ToUpper(job.Text)
			chOut <- job
		}
	}()

	return chOut
}

func saveFiles(chIn <-chan Job) <-chan bool {
	chBool := make(chan bool)

	go func() {
		defer close(chBool)
		for job := range chIn {
			err := os.WriteFile(job.SavePath, []byte(job.Text), 0644)
			if err != nil {
				fmt.Printf("Error writing files: %v", err)
				return
			}
			chBool <- true
		}
	}()

	return chBool
}
