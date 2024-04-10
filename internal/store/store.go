package store

import (
	"context"
	"errors"
	"time"
)

// ErrConflict указывает на конфликт данных в хранилище.
var ErrConflict = errors.New("data conflict")

type Store interface {
	FindRecepient(ctx context.Context, username string) (userID string, err error)
	ListMessages(ctx context.Context, userID string) ([]Message, error)
	GetMessage(ctx context.Context, id int64) (*Message, error)
	SaveMessage(ctx context.Context, userID string, msg Message) error
	// RegisterUser регистрирует нового пользователя
	RegisterUser(ctx context.Context, userID, username string) error
}

// Message описывает объект сообщения
type Message struct {
	ID      int64     // внутренний идентификатор сообщения
	Sender  string    // отправитель
	Time    time.Time // время отправления
	Payload string    // текст сообщения
}
