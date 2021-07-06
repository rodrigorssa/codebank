package usecase

import (
	"time"

	"github.com/rodrigorssa/codebank/domain"
	"github.com/rodrigorssa/codebank/dto"
	"github.com/rodrigorssa/codebank/repository"
)

type UseCaseTransaction struct {
	TransactionRepository repository.TransactionRepository
}

func NewUseCaseTransaction(transactionRepository repository.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{transactionRepository}
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
	return *t, nil
}
