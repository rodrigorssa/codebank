package domain

import (
	"time"

	"github.com/rodrigorssa/codebank/enums"
	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardID string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if t.Amount+creditCard.Balance > creditCard.Limit {
		t.Status = enums.TransactionStatus.REJECTED
	} else {
		t.Status = enums.TransactionStatus.APPROVED
		creditCard.Balance += t.Amount
	}
}
