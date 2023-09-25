package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Mutation struct {
	Id            string    `json:"id"`
	Transaction   int       `json:"transaction"`
	AccountNumber int       `json:"account_number"`
	Balance       int       `json:"balance"`
	TransactionAt time.Time `json:"transaction_at"`
}

func (p *Mutation) Model() interface{} {
	return &UserData{}
}

func (p *Mutation) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	payload["id"] = uuid.New().String()
	payload["transaction_at"] = time.Now()
	payload["code_transaction"] = generateRandomString()
	return payload
}
func (p *Mutation) ModelUpdate(payload map[string]interface{}) map[string]interface{} {
	return payload
}
func (p *Mutation) TbName() string {
	return "mutation_data"
}

func generateRandomString() string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 30
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}
