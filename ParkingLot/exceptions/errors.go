package exceptions

type Error struct {
	statusCode   int
	errorMessage string
	errorCode    string
}

var (
	ErrGeneratingTicketFailure = Error{
		statusCode:   500,
		errorMessage: "some error occured while generating ticket.",
		errorCode:    "PL-001-F",
	}

	ErrGeneratingTicketSuccess = Error{
		statusCode:   200,
		errorMessage: "ticket generated successfully",
		errorCode:    "PL-001-S",
	}
)
