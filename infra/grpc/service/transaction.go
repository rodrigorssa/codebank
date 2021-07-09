package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rodrigorssa/codebank/dto"
	"github.com/rodrigorssa/codebank/enums"
	"github.com/rodrigorssa/codebank/infra/grpc/pb"
	"github.com/rodrigorssa/codebank/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionService struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, in *pb.PaymentRequest) (*empty.Empty, error) {
	transactionDto := dto.Transaction{
		Name:            in.CreditCard.GetName(),
		Number:          in.CreditCard.GetNumber(),
		ExpirationMonth: in.CreditCard.GetExpirationMonth(),
		ExpirationYear:  in.CreditCard.GetExpirationYear(),
		CVV:             in.CreditCard.GetCvv(),
		Amount:          in.GetAmount(),
		Store:           in.GetStore(),
		Description:     in.GetDescription(),
	}
	transaction, err := t.ProcessTransactionUseCase.ProcessTransaction((transactionDto))

	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	if transaction.Status != enums.TransactionStatus.APPROVED {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, enums.TransactionErrors.REJECTED)
	}
	return &empty.Empty{}, nil
}
