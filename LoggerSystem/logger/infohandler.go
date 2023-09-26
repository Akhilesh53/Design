package logger

type InfoHandler struct {
	level         LogType
	next          ILogHandler
	infoObservers []ISinkObserver
}

func getInfoLogHandler(next ILogHandler) *InfoHandler {
	return &InfoHandler{
		level: INFO,
		next:  next,
	}
}

func getDefaultInfoLogHandler() *InfoHandler {
	return &InfoHandler{
		level: INFO,
	}
}

func (h *InfoHandler) SetLevel(level LogType) {
	h.level = level
}

func (h *InfoHandler) SetNext(next ILogHandler) {
	h.next = next
}

func (h *InfoHandler) GetLevel() string {
	return h.level.String()
}

func (h *InfoHandler) GetNext() ILogHandler {
	return h.next
}

func (h *InfoHandler) LogMessage(logType LogType, msg string) {
	if h.level == logType {
		displayMessage(h, msg)
		return
	}
	if h.GetNext() != nil {
		h.GetNext().LogMessage(logType, msg)
	}
}

func (i *InfoHandler) AddSinkObserver(observer ISinkObserver) {
	i.infoObservers = append(i.infoObservers, observer)
}

func (i *InfoHandler) RemoveSinkObserver(observer ISinkObserver) {
	i.infoObservers = append(i.infoObservers, observer)
	newList := []ISinkObserver{}

	for _, obs := range i.infoObservers {
		if obs != observer {
			newList = append(newList, obs)
		}
	}
	i.infoObservers = newList
}

func (i *InfoHandler) NotifyInkObservers(mssg string) {
	for _, obs := range i.infoObservers {
		obs.Update(mssg)
	}
}
