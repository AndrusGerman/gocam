package gocam

// func none() {
// 	ffmpeg.Input("").GetArgs()
// }

// old
// ffmpeg -y -f vfwcap -r 25 -i 0 out.mp4
// ffmpeg -list_devices true -f dshow -i dummy

// get list cams
// pnputil /enum-devices /class Camera /connected

// ffmpeg -f dshow -i video="HP TrueVision HD Camera" out.mp4

// ffmpeg -f dshow -i video="HP TrueVision HD Camera" -vframes 1 out.png^C

// var err = ffmpeg.Input(fmt.Sprintf("video=\"%s\"", cam.cameraName), ffmpeg.KwArgs{
// 	"f": "dshow",
// }).Output("pipe:", ffmpeg.KwArgs{
// 	"vframes": "1",
// }).WithOutput(buffer).ErrorToStdOut().Run()
