package entities

type TransactionType int

const (
	DEBIT TransactionType = iota
	CREDIT
)

func (t TransactionType) String() string {
	switch t {
	case DEBIT:
		return "DEBIT"
	case CREDIT:
		return "CREDIT"
	default:
		return "UNKNOWN"
	}
}

type TransactionStatus int

const (
	PENDING TransactionStatus = iota
	COMPLETED
	CANCELLED
)

func (t TransactionStatus) String() string {
	switch t {
	case PENDING:
		return "PENDING"
	case COMPLETED:
		return "COMPLETED"
	case CANCELLED:
		return "CANCELLED"
	default:
		return "UNKNOWN"
	}
}
