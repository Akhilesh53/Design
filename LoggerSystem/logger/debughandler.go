package logger

type DebugHandler struct {
	Level LogType
	Next  ILogHandler
}

func getDebugLogHandler(next ILogHandler) *DebugHandler {
	return &DebugHandler{
		Level: DEBUG,
		Next:  next,
	}
}

func getDefaultDebugLogHandler() *DebugHandler {
	return &DebugHandler{
		Level: DEBUG,
	}
}

func (h *DebugHandler) SetLevel(level LogType) {
	h.Level = level
}

func (h *DebugHandler) SetNext(next ILogHandler) {
	h.Next = next
}

func (h *DebugHandler) GetLevel() string {
	return h.Level.String()
}

func (h *DebugHandler) GetNext() ILogHandler {
	return h.Next
}

func (h *DebugHandler) LogMessage(logType LogType,msg string) {
	if h.Level == logType {
		displayMessage(h, msg)
		return
	}
	if h.GetNext() != nil {
		h.GetNext().LogMessage(logType,msg)
	}
}
