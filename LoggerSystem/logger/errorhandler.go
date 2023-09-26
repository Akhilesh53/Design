package logger

type ErrorHandler struct {
	level          LogType
	next           ILogHandler
	errorObservers []ISinkObserver
}

func getErrorLogHandler(next ILogHandler) *ErrorHandler {
	return &ErrorHandler{
		level: ERROR,
		next:  next,
	}
}

func getDefaultErrorLogHandler() *ErrorHandler {
	return &ErrorHandler{
		level: ERROR,
	}
}

func (h *ErrorHandler) SetLevel(level LogType) {
	h.level = level
}

func (h *ErrorHandler) SetNext(next ILogHandler) {
	h.next = next
}

func (h *ErrorHandler) GetLevel() string {
	return h.level.String()
}

func (h *ErrorHandler) GetNext() ILogHandler {
	return h.next
}

func (h *ErrorHandler) LogMessage(logType LogType, msg string) {
	if h.level == logType {
		displayMessage(h, msg)
		return
	}
	if h.GetNext() != nil {
		h.GetNext().LogMessage(logType, msg)
	}
}

func (i *ErrorHandler) AddSinkObserver(observer ISinkObserver) {
	i.errorObservers = append(i.errorObservers, observer)
}

func (i *ErrorHandler) RemoveSinkObserver(observer ISinkObserver) {
	i.errorObservers = append(i.errorObservers, observer)
	newList := []ISinkObserver{}

	for _, obs := range i.errorObservers {
		if obs != observer {
			newList = append(newList, obs)
		}
	}
	i.errorObservers = newList
}

func (i *ErrorHandler) NotifyInkObservers(mssg string) {
	for _, obs := range i.errorObservers {
		obs.Update(mssg)
	}
}
