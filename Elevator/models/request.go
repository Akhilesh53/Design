package models

import "pattern/Elevator/enums"

type Request struct {
	sourceFloor uint32
	destFloor uint32
	direction enums.Direction
	requestType enums.RequestType
}