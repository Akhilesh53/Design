package enums

type SpotStatus int

const (
	AVAILABLE SpotStatus = iota
	OCCUPIED
	NOT_AVAILABLE
)
