package raw

type Header struct {
	FileIdentifier string
	Metadata       Metadata
	ImageWidth     int
	ImageHeight    int
}

type Metadata struct {
	CameraModel     string
	ExposureTime    float64
	ISO             int
	Aperture        float64
	Timestamp       string
	WhiteBalance    string
	LensInformation string
	FocusInfo       FocusInformation
}

type FocusInformation struct {
	FocusPoints   []FocusPoint
	FocusDistance float64
}

type FocusPoint struct {
	X int
	Y int
}

type RAW struct {
	Header     Header
	SensorData SensorData
}

type SensorData struct {
	PixelData    [][]int
	BitDepth     int
	BayerPattern BayerPattern
}

type BayerPattern struct {
	Pattern [][]string
}

// BayerFilter to rgb image

// Bayerfilter to yuv image
