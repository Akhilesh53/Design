package enums

type PaymentMode int

const (
	CASH PaymentMode = iota
	CREDIT_CARD
	DEBIT_CARD
	UPI
)
