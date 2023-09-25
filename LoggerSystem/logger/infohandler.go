package logger

type InfoHandler struct {
	Level LogType
	Next  ILogHandler
}

func getInfoLogHandler(next ILogHandler) *InfoHandler {
	return &InfoHandler{
		Level: INFO,
		Next:  next,
	}
}

func getDefaultInfoLogHandler() *InfoHandler {
	return &InfoHandler{
		Level: INFO,
	}
}

func (h *InfoHandler) SetLevel(level LogType) {
	h.Level = level
}

func (h *InfoHandler) SetNext(next ILogHandler) {
	h.Next = next
}

func (h *InfoHandler) GetLevel() string {
	return h.Level.String()
}

func (h *InfoHandler) GetNext() ILogHandler {
	return h.Next
}

func (h *InfoHandler) LogMessage(logType LogType, msg string) {
	if h.Level == logType {
		displayMessage(h, msg)
		return
	}
	if h.GetNext() != nil {
		h.GetNext().LogMessage(logType, msg)
	}
}
