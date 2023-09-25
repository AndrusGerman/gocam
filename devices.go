package gocam

import (
	"bufio"
	"os/exec"
	"strings"
)

const (
	DeviceDescriptionKey = "Device Description"
)

func GetListDevices() ([]string, error) {
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
		if strings.Contains(text, DeviceDescriptionKey) {
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
