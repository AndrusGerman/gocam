package gocam

func NewCamera(cameraName string) (*Camera, error) {
	return &Camera{
		cameraName: cameraName,
	}, nil
}

type Camera struct {
	cameraName string
}
