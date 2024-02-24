package entities

type Meeting struct {
	meetingId     uint8
	meetingRoomId uint8
	description   string
	participants  []INotificationObserver
	host          INotificationObserver
}

func NewMeeting(description string, host INotificationObserver, id, roomId uint8) *Meeting {
	return &Meeting{
		description:   description,
		meetingId:     id,
		meetingRoomId: roomId,
		participants:  make([]INotificationObserver, 0),
		host:          host,
	}
}

func (m *Meeting) NotifyUsers() {
	for _, attendee := range m.participants {
		attendee.SendNotificationToUser(m.description)
	}
	m.host.SendNotificationToUser(m.description)
}
