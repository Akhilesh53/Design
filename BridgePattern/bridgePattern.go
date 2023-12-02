package main

import (
	"fmt"
	"pattern/BridgePattern/entities"
)

func main() {
	tvDevice, err := GetDevice("TV")
	if err != nil {
		fmt.Println(err)
		return
	}
	clientACRemote := entities.NewACRemote(tvDevice)

	clientACRemote.TogglePower()
}
