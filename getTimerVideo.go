package gocam

import (
	"bytes"
	"fmt"
	"math"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// TODO: is very basic
func (cam *Camera) GetTimeVideo(time time.Duration) ([]byte, error) {
	var buffer = new(bytes.Buffer)
	var seconds = int(math.Round(time.Seconds()))
	var err = ffmpeg.Input(fmt.Sprintf("video=%s", cam.cameraName), ffmpeg.KwArgs{
		"f":  "dshow",
		"ss": "00:00:00",
		"to": fmt.Sprintf(
			"00:00:%d",
			seconds,
		),
	}).Output("pipe:", ffmpeg.KwArgs{
		"c:v": "libx264",
		"f":   "flv",
	}).WithOutput(buffer).Run()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// ffmpeg -f dshow -ss 00:00:00 -to 00:02:00 -i video="HP TrueVision HD Camera" out.mp4
// ffmpeg -f dshow -ss 00:00:00 -to 00:00:01 -i video="HP TrueVision HD Camera" -c:v libvpx-vp9 -f webm  pipe:1 > x.webm

//ffmpeg -f dshow -ss 00:00:00 -to 00:00:01 -i video="HP TrueVision HD Camera" -c:v libx264 -f flv  pipe:1 > x.mp4
//ffmpeg -f dshow -ss 00:00:00 -to 00:00:10 -i video="HP TrueVision HD Camera" -c:v libx264 -f flv  a.mp4
