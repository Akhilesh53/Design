package entities

type MeetingRoom struct {
	roomId   uint8
	calendar *Calendar
	capacity uint8
}

func NewMeetingRoom(roomId, capacity uint8) *MeetingRoom {
	return &MeetingRoom{
		roomId:   roomId,
		capacity: capacity,
	}
}

func (r *MeetingRoom) AddCalendar(calendar *Calendar) *MeetingRoom {
	r.calendar = calendar
	return r
}

func (r *MeetingRoom) IsAvailable(slot *TimeSlot) bool {
	return r.calendar.CheckAvailability(slot)
}

func (r *MeetingRoom) ScheduleMeeting(slot *TimeSlot, meeting *Meeting) {
	// schedule meeting on calendar
}

func (r *MeetingRoom) GetRoomId() uint8 {
	return r.roomId
}
