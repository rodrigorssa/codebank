package enums

type TransactionStatusEnum struct {
	REJECTED string
	APPROVED string
}

type TransactionErrorsEnum struct {
	NOT_FOUND string
}

var TransactionStatus = TransactionStatusEnum{
	REJECTED: "REJECTED",
	APPROVED: "APPROVED",
}

var TransactionErrors = TransactionErrorsEnum{
	NOT_FOUND: "Credit card does not exist!",
}
