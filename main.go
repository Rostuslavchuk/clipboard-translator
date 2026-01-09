package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	gde "github.com/joho/godotenv"
	"log"
	"os"
)

type EnvData struct {
	apiKey string
	url    string
}

func main() {
	if err := gde.Load(); err != nil {
		log.Fatal(err)
	}

	envData := EnvData{
		apiKey: os.Getenv("API_KEY"),
		url:    os.Getenv("URL"),
	}

	var target_lang string

	fmt.Print("Input target language like: en, uk, es: ")
	fmt.Scanf("%s", &target_lang)

	if target_lang == "" {
		log.Fatal("Target language must not be empty!")
	}

	var lastClipboard string
	for {
		text, err := clipboard.ReadAll()
		if err != nil {
			log.Fatal(err)
			continue
		}
		if text == lastClipboard || text == "" {
			continue
		}

		if text != lastClipboard {
			translated, err := TranslateText(text, target_lang, envData.apiKey, envData.url)
			if err != nil {
				log.Print(err)
				continue
			}

			clipboard.WriteAll(translated)
			lastClipboard = text
		}
	}
}
