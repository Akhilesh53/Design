package main

import (
	"fmt"
	"pattern/MeetingScheduler/entities"
	"time"
)

func main() {

	meetingScheduler := entities.NewMeetingScheduler(20, entities.NewFCFSRoomBookingStrategy())
	// create room
	room1 := entities.NewMeetingRoom(1, 100)

	// create a calendar
	calendar := entities.NewCalendar(1)
	room1.AddCalendar(calendar)

	// add slot to calendar
	slot := entities.NewTimeSlot("2024-02-24", "2024-02-24", time.Now(), 15*time.Minute)
	calendar.AddSlot(slot)

	meetingScheduler.AddMeetingRoom(room1)

	// schedule a meeting
	meetingHost := entities.NewUser(456, "Akhilesh", "ak.m@gmail.com", "990xxx")
	meeting, err := meetingScheduler.ScheduleMeeting(slot, meetingHost)

	if err != nil {
		fmt.Println(err)
		return
	}

	// send an additional template in notification
	meetingScheduler.Notify(meeting)
}
