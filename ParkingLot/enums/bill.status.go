package enums

type BillStatus int

const (
	SUCCESSFUL BillStatus = iota
	PENDING
	UNSUCCESSFUL
)
