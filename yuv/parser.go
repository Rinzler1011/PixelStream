package yuv

import (
	"fmt"
	"io"
	"os"
)

type YUV struct {
	Width  int
	Height int
	YSize  int
	USize  int
	VSize  int
	Y      []byte
	U      []byte
	V      []byte
}

type RGB struct {
	Width       int
	Height      int
	ChannelSize int
	Pixels      []Pixel
}

type Pixel struct {
	R, G, B uint8
}

// RGB to YUV function

func ReadRGBFile(filepath string) ([]Pixel, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	var pixels []Pixel
	buffer := make([]byte, 3)

	for {
		_, err := io.ReadFull(file, buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}

		pixel := Pixel{
			R: buffer[0],
			G: buffer[1],
			B: buffer[2],
		}

		pixels = append(pixels, pixel)
	}

	return pixels, nil

}

func RGBToYUV(pixels []Pixel, width int, height int, samplingRate int) (YUV, error) {

	var ySize, uSize, vSize int

	switch {
	case samplingRate == 444:
		ySize = width * height
		uSize = width * height
		vSize = width * height
	case samplingRate == 422:
		ySize = width * height
		uSize = (width / 2) * height
		vSize = (width / 2) * height
	case samplingRate == 420:
		ySize = width * height
		uSize = (width / 2) * (height / 2)
		vSize = (width / 2) * (height / 2)
	case samplingRate == 111:
		ySize = width * height
		uSize = (width / 4) * (height / 4)
		vSize = (width / 4) * (height / 4)
	default:
		return YUV{}, fmt.Errorf("error: Unsupported sampling rate: %d", samplingRate)
	}

	yuv := YUV{
		Width:  width,
		Height: height,
		YSize:  ySize,
		USize:  uSize,
		VSize:  vSize,
		Y:      make([]byte, ySize),
		U:      make([]byte, uSize),
		V:      make([]byte, vSize),
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			yIndex := i*width + j
			Y, U, V := RGBToBT2020(444, pixels[yIndex].R, pixels[yIndex].G, pixels[yIndex].B)

			yuv.Y[yIndex] = byte(Y)

			if samplingRate == 444 {
				yuv.U[yIndex] = byte(U)
				yuv.V[yIndex] = byte(V)
			} else if samplingRate == 422 {
				if j%2 == 0 {
					uIndex := (i*width + j) / 2
					yuv.U[uIndex] = byte(U)
					yuv.V[uIndex] = byte(V)
				}
			} else if samplingRate == 420 {
				if i%2 == 0 && j%2 == 0 {
					uIndex := (i/2)*(width/2) + (j / 2)
					yuv.U[uIndex] = byte(U)
					yuv.V[uIndex] = byte(V)
				}
			} else if samplingRate == 111 {
				if i%4 == 0 && j%4 == 0 {
					uIndex := (i/4)*(width/4) + (j / 4)
					yuv.U[uIndex] = byte(U)
					yuv.V[uIndex] = byte(V)
				}
			}
		}
	}

	return yuv, nil
}

// Limit Function to keep values

// RGB to BT.2020 (Rec.2020) YUV | Conversion algorithm ---------- Work on doing calcs only on the samplingRate

func RGBToBT2020(samplingRate int, R uint8, G uint8, B uint8) (Y float64, U float64, V float64) {

	// BT.2020 (Rec.2020) - UHD / wide-gamut
	// Y = 0.2627 * R + 0.6780 * G + 0.0593 * B
	// U = -0.13963 * R - 0.36037 * G + 0.5 * B + 128
	// V = 0.5 * R - 0.45979 * G - 0.04021 * B + 128

	if samplingRate == 444 {
		Y = (0.2627 * float64(R)) + (0.6780 * float64(G)) + (0.0593 * float64(B))
		U = (-0.13963 * float64(R)) - (0.36037 * float64(G)) + (0.5 * float64(B)) + 128
		V = (0.5 * float64(R)) - (0.45979 * float64(G)) - (0.04021 * float64(B)) + 128

	}
	return Y, U, V

}

// YUV to RGB function

// RAW to YUV - 8bit

// RAW to YUV - 10bit
