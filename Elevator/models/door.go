package models

import "pattern/Elevator/enums"

type Door struct {
	doorId     uint32
	doorStatus enums.Status
}
