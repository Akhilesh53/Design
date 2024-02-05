package enums

type RequestType int

const (
	INTERNAL_REQUEST RequestType = iota
	EXTERNAL_REQUEST
)
