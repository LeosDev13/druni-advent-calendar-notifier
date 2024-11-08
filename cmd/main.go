package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	resp, err := http.Get("https://www.druni.es/calendario-adviento-druni-24-dias")
	if err != nil {
		panic("handle error")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	bodyString := string(body)

	match, _ := regexp.MatchString(`<span>No\s*disponible</span>`, bodyString)

	if match {
		log.Println("seguimos sin suerte :(")
	} else {
		sendMessage("Entra en druni rápido que se acaban los calendarios https://www.druni.es/calendario-adviento-druni-24-dias")
	}
}

func sendMessage(text string) (bool, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", os.Getenv("TOKEN"), os.Getenv("CHAT_ID"), text)
	body, _ := json.Marshal(map[string]string{
		"text":    text,
		"chat_id": os.Getenv("CHAT_ID"),
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	log.Println("Message %s was sent", text)
	log.Println("Response JSON: %s", string(body))

	return true, nil
}
