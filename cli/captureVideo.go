package main

import (
	"os"
	"time"

	"github.com/AndrusGerman/gocam"
)

func captureVideo() {
	devices, err := gocam.GetListDevices()
	if err != nil {
		panic(err)
	}

	camera, err := gocam.NewCamera(devices[0])
	if err != nil {
		panic(err)
	}

	bt, err := camera.GetTimeVideo(time.Second * 120)
	if err != nil {
		panic(err)
	}

	os.WriteFile("video.mp4", bt, 0666)
}
