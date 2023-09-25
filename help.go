package gocam

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	// err := ffmpeg.Input("./a.mp4").
	// 	Output("./out1.mp4", ffmpeg.KwArgs{"c:v": "libx265"}).
	// 	OverWriteOutput().ErrorToStdOut().Run()
}

func none() {
	ffmpeg.Input("")
}

// old
// ffmpeg -y -f vfwcap -r 25 -i 0 out.mp4
// ffmpeg -list_devices true -f dshow -i dummy

// get list cams
// pnputil /enum-devices /class Camera /connected

// ffmpeg -f dshow -i video="HP TrueVision HD Camera" out.mp4

// ffmpeg -f dshow -i video="HP TrueVision HD Camera" -vframes 1 out.png^C
