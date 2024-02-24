package entities

type Calendar struct {
	scheduledMeetings []*Meeting
	timeSlots         []*TimeSlot
	meetingRoomId     uint8
}

func NewCalendar(meetingRoomId uint8) *Calendar {
	return &Calendar{
		meetingRoomId:     meetingRoomId,
		scheduledMeetings: make([]*Meeting, 0),
		timeSlots:         make([]*TimeSlot, 0),
	}
}

func (c *Calendar) AddSlot(slot *TimeSlot) *Calendar {
	c.timeSlots = append(c.timeSlots, slot)
	return c
}

func (c *Calendar) CheckAvailability(slot *TimeSlot) bool {
	// logic will be changed
	for _, availSlot := range c.timeSlots {
		if availSlot.GetStartTime().Equal(slot.startTime) && availSlot.GetEndTime().Equal(slot.startTime.Add(slot.duration)) {
			return true
		}
	}
	return false
}
