package gocam

import (
	"fmt"
	"image"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func (cam *Camera) GetPhoto() image.Image {
	ffmpeg.Input(fmt.Sprintf("video=\"%s\"", cam.cameraName), ffmpeg.KwArgs{
		"f": "dshow",
	}).Output("out.png", ffmpeg.KwArgs{
		"vframes": "1",
	}).Run()
	return nil
}
