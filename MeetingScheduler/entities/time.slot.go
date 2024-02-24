package entities

import (
	"time"
)

type TimeSlot struct {
	startDate string
	endDate   string
	startTime time.Time
	duration  time.Duration
}

func NewTimeSlot(startDate string, endDate string, startTime time.Time, duration time.Duration) *TimeSlot {
	return &TimeSlot{startDate: startDate, endDate: endDate, startTime: startTime, duration: duration}
}

func (slot *TimeSlot) GetStartTime() time.Time {
	return slot.startTime
}

func (slot *TimeSlot) GetEndTime() time.Time {
	return slot.startTime.Add(slot.duration)
}
