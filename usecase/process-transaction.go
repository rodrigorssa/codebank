package usecase

import (
	"encoding/json"
	"time"

	"github.com/rodrigorssa/codebank/domain"
	"github.com/rodrigorssa/codebank/dto"
	"github.com/rodrigorssa/codebank/enums"
	"github.com/rodrigorssa/codebank/infra"
	"github.com/rodrigorssa/codebank/repository"
)

type UseCaseTransaction struct {
	TransactionRepository repository.TransactionRepository
	KafkaProducer         infra.KafkaProducer
}

func NewUseCaseTransaction(transactionRepository repository.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV
	ccBalenceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = ccBalenceAndLimit.ID
	creditCard.Limit = ccBalenceAndLimit.Limit
	creditCard.Balance = ccBalenceAndLimit.Balance

	t := domain.NewTransaction()
	t.CreditCardID = creditCard.ID
	t.Amount = transactionDto.Amount
	t.Store = transactionDto.Store
	t.Description = transactionDto.Description
	t.CreatedAt = time.Now()

	transactionDto.ID = t.ID
	transactionDto.CreatedAt = t.CreatedAt

	transactionJson, err := json.Marshal(transactionDto)

	if err != nil {
		return domain.Transaction{}, err
	}

	err = u.KafkaProducer.Publish(string(transactionJson), enums.PAYMENT_TOPIC)

	if err != nil {
		return domain.Transaction{}, err
	}

	return *t, nil
}
