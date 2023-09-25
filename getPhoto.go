package gocam

import (
	"bytes"
	"fmt"
	"image/png"

	ffmpeg "github.com/u2takey/ffmpeg-go"

	"image"
)

func (cam *Camera) GetPhoto() (image.Image, error) {
	var buffer = new(bytes.Buffer)
	var err = ffmpeg.Input(fmt.Sprintf("video=%s", cam.cameraName), ffmpeg.KwArgs{
		"f": "dshow",
	}).Output("pipe:", ffmpeg.KwArgs{
		"vframes": "1",
		"format":  "image2",
		"vcodec":  "png",
	}).WithOutput(buffer).Run()
	if err != nil {
		return nil, err
	}
	return png.Decode(buffer)
}
