package entities

type INotificationObserver interface {
	SendNotificationToUser(string)
}
