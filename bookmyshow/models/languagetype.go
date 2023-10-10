package models

type Language int

const (
	HINDI Language = iota
	ENGLISH
	PUNJABI
)

func (l Language) String() string {
	switch l {
	case HINDI:
		return "Hindi"
	case ENGLISH:
		return "English"
	case PUNJABI:
		return "Punjabi"
	default:
		return "Unknown"
	}
}
