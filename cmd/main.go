package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const druniAdventCalendarURL = "https://www.druni.es/calendario-adviento-druni-24-dias"
const telegramBotUrl = "https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s"

var (
	notAvailableRegex   = regexp.MustCompile(`<span>No\s*disponible</span>`)
	productSoldOutRegex = regexp.MustCompile(`<p class="sold-out">Producto\s*agotado</p>`)
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	resp, err := http.Get(druniAdventCalendarURL)
	if err != nil {
		log.Printf("Error occurred: %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK HTTP status: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	bodyString := string(body)

	if shouldSendMessage(bodyString) {
		sendMessage(fmt.Sprintf("Entra en druni rápido que se acaban los calendarios %s", druniAdventCalendarURL))
	}
}

func shouldSendMessage(text string) bool {
	notAvailableMatch := notAvailableRegex.MatchString(text)
	productSoldOutMatch := productSoldOutRegex.MatchString(text)

	if !notAvailableMatch && !productSoldOutMatch {
		return true
	}
	return false
}

func sendMessage(text string) {
	token := os.Getenv("TOKEN")
	chatID := os.Getenv("CHAT_ID")
	if token == "" || chatID == "" {
		log.Fatal("TOKEN or CHAT_ID is missing in the environment variables.")
	}

	url := fmt.Sprintf(telegramBotUrl, os.Getenv("TOKEN"), os.Getenv("CHAT_ID"), text)
	body, _ := json.Marshal(map[string]string{
		"text":    text,
		"chat_id": os.Getenv("CHAT_ID"),
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Printf("Error occurred: %v", err)
		return
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error occurred: %v", err)
		return
	}

	log.Printf("Message %s was sent", text)
	log.Printf("Response JSON: %s", string(body))
}
