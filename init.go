package gocam

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func init() {
	ffmpeg.LogCompiledCommand = false
}
