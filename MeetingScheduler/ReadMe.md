Design Meeting Scheduler. 

Given (n) meeting rooms. Users can book any meeting room from (start time to end time) if the room is available. 
If a room is not available, it should be shown at that moment only.
All the participants of a meeting should receive a notification related to the same.
Use the calendar for tracking dates and time.
Store the history of all the meetings scheduled.

Room Assignment (different strategies)
User availability (how to check which user is available at what time)
What kind of meetings will be there (recurrence or one time)
Capacity of meeting rooms
Calendar can be specific to user profile and meeting rooms. which to consider.

--------------------------------
Entities
--------------------------------
1) MeetingRoom (roomid, calendar, capacity)
2) MeetingScheduler (List<Meeting>historicalMeetings, List<MeetingRoom>)
3) Meeting (metingid, MeetingRoom, description, List<User>)
4) Calendar (MeetingRoom, map[date]List<TimeSlots>, map[string]*Meeting)
5) User (Host, Attendees) (id, name, email, who) 
6) Notification (Id,Meeting)
7) RoomAssignmentStrategy <<>>
8) TimeSlots (startdate, enddate, starttime, endtime)

Detailed Description of Entities
-----------------------------------

// this itself will act as meeting observable <<>>
Meeting{
    meetingId
    MeetingRoom
    Description
    List<Attendee> participants
    Host
}

CreateMeeting(meetingRoomId int, description string, List<Attendees> users, host User) meetingId
CancelMeeting(meetingId int)
UpdateMeetingDescription(meetingId int, description string)
AddAttendeestoMeeting(meetingId, List<Attendees> attendees)
RemoveAttendeestoMeeting(meetingId, Attendee attendee)
// Getter Setter for all tags

MeetingRoom{
    roomId
    calendar
    capacity
}

IsAvailable(TimeSlot slot) bool
scheduleMeeting(TimeSlot slot) meetingId

Calendar{
    List<TimeSlot> slots
    MeetingRoom
    List<Meeting>ScheduledEvents
}

checkAvailability(TimeSlot slot) bool 
scheduleNewMeetings(TimeSlot slot) meetingId

User{
    userId
    name
    email
    contact
}

Host{
    User
}

Atendees{
    User
}

TimeSlot{
    startDate
    endDate
    startTime
    endTime
}

Notification{
    notifId
    message
    meeting
    List<User> participants
}

NotifificationObserver <<>>              -- User
SendNotification(MeetingObservable)


MeetingScheduler{
    MAX_HOSTORY_MEETINGS 
    List<Meeting> historicalMeetings
    List<MeetingRoom> meetingsRooms
    NotificationObserver notificationObserver
}


----------------------------------------------------------------
Design Pattern
----------------------------------------------------------------
Observer Design Pattern for Notification Sending 

----------------------------------------------------------------
Relation Between Entities
----------------------------------------------------------------
1) Meeting and Meeting Room
    1 : M
    M : M
    M : M

2) Meeting and Attendees 
    1 : M
    M : 1
    M : M

2) Meeting and Host
    1 : M
    M : 1
    M : M

4) Meeting Room and Calendar
    1 : 1
    1 : 1
    1 : 1 

5) Calendar and Meeting
    1 : M
    1 : 1
    1 : M


