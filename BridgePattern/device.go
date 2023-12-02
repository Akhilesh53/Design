package main

import (
	"fmt"
	"pattern/BridgePattern/entities"
	"pattern/BridgePattern/interfaces"
)

func GetDevice(deviceType string) (interfaces.Device, error) {
	switch deviceType {
	case "AC":
		return entities.NewACDevice(), nil
	default:
		return nil, fmt.Errorf("device type not supported: %s", deviceType)
	}
}
