package domain

type Device struct {
	UA string
	IP string

	OS        string
	OSVersion string
	Make      string
	Model     string

	DeviceType int
}
