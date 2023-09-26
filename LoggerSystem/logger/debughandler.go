package logger

type DebugHandler struct {
	level          LogType
	next           ILogHandler
	debugObservers []ISinkObserver
}

func getDebugLogHandler(next ILogHandler) *DebugHandler {
	return &DebugHandler{
		level: DEBUG,
		next:  next,
	}
}

func getDefaultDebugLogHandler() *DebugHandler {
	return &DebugHandler{
		level: DEBUG,
	}
}

func (h *DebugHandler) SetLevel(level LogType) {
	h.level = level
}

func (h *DebugHandler) SetNext(next ILogHandler) {
	h.next = next
}

func (h *DebugHandler) GetLevel() string {
	return h.level.String()
}

func (h *DebugHandler) GetNext() ILogHandler {
	return h.next
}

func (h *DebugHandler) LogMessage(logType LogType, msg string) {
	if h.level == logType {
		displayMessage(h, msg)
		return
	}
	if h.GetNext() != nil {
		h.GetNext().LogMessage(logType, msg)
	}
}

func (i *DebugHandler) AddObserver(observer ISinkObserver) {
	i.debugObservers = append(i.debugObservers, observer)
}

func (i *DebugHandler) RemoveObserver(observer ISinkObserver) {
	i.debugObservers = append(i.debugObservers, observer)
	newList := []ISinkObserver{}

	for _, obs := range i.debugObservers {
		if obs != observer {
			newList = append(newList, obs)
		}
	}
	i.debugObservers = newList
}

func (i *DebugHandler) Notify(mssg string) {
	for _, obs := range i.debugObservers {
		obs.Update(mssg)
	}
}
