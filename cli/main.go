package main

import (
	"github.com/AndrusGerman/gocam"
)

func main() {
	devices, err := gocam.GetListDevices()
	if err != nil {
		panic(err)
	}

	camera, err := gocam.NewCamera(devices[0])
	if err != nil {
		panic(err)
	}

	camera.GetPhoto()
}
