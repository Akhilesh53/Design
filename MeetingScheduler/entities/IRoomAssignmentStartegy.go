package entities

type IRoomAssignmentStrategy interface {
	BookRoom(rroms []*MeetingRoom, slot *TimeSlot) *MeetingRoom
}

type FCFSRoomBookingStrategy struct {
}

func NewFCFSRoomBookingStrategy() *FCFSRoomBookingStrategy {
	return &FCFSRoomBookingStrategy{}
}

func (s *FCFSRoomBookingStrategy) BookRoom(rooms []*MeetingRoom, slot *TimeSlot) *MeetingRoom {
	for _, room := range rooms {
		if room.IsAvailable(slot) {
			return room
		}
	}
	return nil
}
