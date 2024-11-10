package repository

type SenderRepository interface {
	SendMessage(id, text string) error
}
