package models

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type UserData struct {
	Id            string    `json:"id"`
	AccountNumber string    `json:"account_number"`
	Balance       int       `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type Users struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	NIK         int    `json:"nik"`
	PhoneNumber string `json:"phone_number"`
	PinNumber   int    `json:"pin_number"`
	UserData
}

func (p *Users) Model() interface{} {
	return &Users{}
}

func (p *Users) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	payload["id"] = uuid.New().String()
	payload["balance"] = 0
	payload["account_number"] = p.GenerateAccountNumber()
	payload["created_at"] = time.Now()
	payload["updated_at"] = time.Now()
	return payload
}
func (p *Users) ModelUpdate(payload map[string]interface{}) map[string]interface{} {
	if payload["account_number"] != "" {
		payload["updated_at"] = time.Now()
		return payload
	}
	return nil
}
func (p *Users) TbName() string {
	return "users"
}

func (p *UserData) GenerateAccountNumber() int {
	codeBank := 3130

	min := 10000000
	max := 99999999
	randomNumber := rand.Intn(max-min) + min
	strAccNumber := strconv.Itoa(codeBank) + strconv.Itoa(randomNumber)
	accNumber, _ := strconv.Atoi(strAccNumber)

	return accNumber
}
