package gocam

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const (
	pnpDeviceDescriptionKey = "Device Description"
	dshowVideoKey           = "(video)"
)

type GetListDevicesConfiguration struct {
	Pnputil bool
}

func GetListDevices(cfg ...GetListDevicesConfiguration) ([]string, error) {
	var configuration GetListDevicesConfiguration
	if len(cfg) != 0 {
		configuration = cfg[0]
	}
	if configuration.Pnputil {
		return getListDevicesPnpUtil(configuration)
	}
	return getListDevicesFfmpeg(configuration)
}

func getListDevicesPnpUtil(configuration GetListDevicesConfiguration) ([]string, error) {
	// get list devices info
	cmd := exec.Command("pnputil", "/enum-devices", "/class", "Camera", "/connected")
	bt, err := cmd.Output()
	if err != nil {
		return make([]string, 0), err
	}

	// filter Device Descrition
	var listDeviceDescription []string
	scanner := bufio.NewScanner(strings.NewReader(string(bt)))
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, pnpDeviceDescriptionKey) {
			listDeviceDescription = append(listDeviceDescription, text)
		}
	}

	// Get Device Description Name
	var listDevices = make([]string, len(listDeviceDescription))
	for index, deviceDescription := range listDeviceDescription {
		indexCut := strings.Index(deviceDescription, ":")
		listDevices[index] = strings.TrimSpace(deviceDescription[indexCut+1:])
	}

	return listDevices, nil
}

func getListDevicesFfmpeg(configuration GetListDevicesConfiguration) ([]string, error) {
	var stderr = new(bytes.Buffer)
	var stdout = new(bytes.Buffer)
	err := ffmpeg.Input("dummy", ffmpeg.KwArgs{
		"f":            "dshow",
		"list_devices": "true",
	}).WithErrorOutput(stderr).WithOutput(stdout).Run()

	if !strings.Contains(stderr.String(), "dummy") {
		return nil, err
	}

	var videoDevicesLineRaw []string
	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, dshowVideoKey) {
			videoDevicesLineRaw = append(videoDevicesLineRaw, text)
		}
	}

	var listDevices = make([]string, len(videoDevicesLineRaw))

	for i, v := range videoDevicesLineRaw {
		idx := strings.Index(v, `"`)
		cutStr := v[idx+1:]
		idx2 := strings.Index(cutStr, `"`)
		cutStr = cutStr[:idx2]
		listDevices[i] = cutStr
	}
	return listDevices, nil
}
