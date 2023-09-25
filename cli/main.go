package main

import (
	"fmt"

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

	photo, err := camera.GetPhoto()
	if err != nil {
		panic(err)
	}
	fmt.Println("Photo: ", photo)
}
