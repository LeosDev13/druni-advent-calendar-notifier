package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/leosdev13/druni-scrapper/pkg/repository"
)

type TelegramSenderRepository struct {
	BotToken string
}

func NewTelegramRepository(botToken string) repository.SenderRepository {
	return &TelegramSenderRepository{
		BotToken: botToken,
	}
}

const telegramBotUrl = "https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s"

func (t *TelegramSenderRepository) SendMessage(id, text string) error {
	url := fmt.Sprintf(telegramBotUrl, t.BotToken, id, text)

	body, err := json.Marshal(map[string]string{
		"chat_id": id,
		"text":    text,
	})
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, HTTP status: %d", resp.StatusCode)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	return nil
}
