package entities

import "errors"

type MeetingScheduler struct {
	maxHistoryMeetings  uint8
	meetingRooms        []*MeetingRoom
	historicalMeetings  []*Meeting
	roomBookingStrategy IRoomAssignmentStrategy
}

func NewMeetingScheduler(maxHistoryMeetings uint8, strategy IRoomAssignmentStrategy) *MeetingScheduler {
	return &MeetingScheduler{
		maxHistoryMeetings:  maxHistoryMeetings,
		meetingRooms:        make([]*MeetingRoom, 0),
		historicalMeetings:  make([]*Meeting, 0),
		roomBookingStrategy: strategy,
	}
}

func (ms *MeetingScheduler) Notify(meeting IMeetingObservable) {
	meeting.NotifyUsers()
}

func (ms *MeetingScheduler) AddMeetingRoom(room *MeetingRoom) *MeetingScheduler {
	ms.meetingRooms = append(ms.meetingRooms, room)
	return ms
}

func (ms *MeetingScheduler) ScheduleMeeting(slot *TimeSlot, host INotificationObserver) (*Meeting, error) {
	room := ms.roomBookingStrategy.BookRoom(ms.meetingRooms, slot)
	if room != nil {
		meeting := NewMeeting("test meeting", host, 78, room.GetRoomId())
		room.ScheduleMeeting(slot, meeting)
		return meeting, nil
	}
	return nil, errors.New("no slot found in any room")
}

func (ms *MeetingScheduler) ScheduleMeetingInRoom(roomId uint8, slot *TimeSlot) (*Meeting, error) {
	return nil, errors.New("no slot found in any room")
}
