# GoCam
In simple words it is a tool that allows you to take photos on windows with golang.
Yes, this returns an image.Image.
And of course I plan to make this work on Linux. The goal is to have a simple way to take a photo on linux and windows with golang
### Get Photo
```go
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

```