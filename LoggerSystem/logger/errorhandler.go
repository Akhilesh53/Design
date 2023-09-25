package logger

type ErrorHandler struct {
	Level LogType
	Next  ILogHandler
}

func getErrorLogHandler(next ILogHandler) *ErrorHandler {
	return &ErrorHandler{
		Level: ERROR,
		Next:  next,
	}
}

func getDefaultErrorLogHandler() *ErrorHandler {
	return &ErrorHandler{
		Level: ERROR,
	}
}

func (h *ErrorHandler) SetLevel(level LogType) {
	h.Level = level
}

func (h *ErrorHandler) SetNext(next ILogHandler) {
	h.Next = next
}

func (h *ErrorHandler) GetLevel() string {
	return h.Level.String()
}

func (h *ErrorHandler) GetNext() ILogHandler {
	return h.Next
}

func (h *ErrorHandler) LogMessage(logType LogType, msg string) {
	if h.Level == logType {
		displayMessage(h, msg)
		return
	}
	if h.GetNext() != nil {
		h.GetNext().LogMessage(logType, msg)
	}
}
