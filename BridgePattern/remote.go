package main

import (
	"fmt"
	"pattern/BridgePattern/entities"
	"pattern/BridgePattern/interfaces"
)

func GetRemote(remoteType string, device interfaces.Device) (interfaces.Remote, error) {
	switch remoteType {
	case "AC":
		return entities.NewACRemote(device), nil
	default:
		return nil, fmt.Errorf("remote type not supported %s", remoteType)
	}
}
