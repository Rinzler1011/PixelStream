package jpeg

import (
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

// 50% quality - increase qaulity half every value in quantization table
var LumQuantizationTable = [8][8]int{
	{16, 11, 10, 16, 24, 40, 51, 61},
	{12, 12, 14, 19, 26, 58, 60, 55},
	{14, 13, 16, 24, 40, 57, 69, 56},
	{14, 17, 22, 29, 51, 87, 80, 62},
	{18, 22, 37, 56, 68, 109, 103, 77},
	{24, 35, 55, 64, 81, 104, 113, 92},
	{49, 64, 78, 87, 103, 121, 120, 101},
	{72, 92, 95, 98, 112, 100, 103, 99},
}

var ChromaQuantizationTable = [8][8]int{
	{17, 18, 24, 47, 99, 99, 99, 99},
	{18, 21, 26, 66, 99, 99, 99, 99},
	{24, 26, 56, 99, 99, 99, 99, 99},
	{47, 66, 99, 99, 99, 99, 99, 99},
	{99, 99, 99, 99, 99, 99, 99, 99},
	{99, 99, 99, 99, 99, 99, 99, 99},
	{99, 99, 99, 99, 99, 99, 99, 99},
	{99, 99, 99, 99, 99, 99, 99, 99},
}

func ReadYUV(filepath string, width int, height int) {
	file, err := os.Open(filepath)

	if err != nil {
		return // error code
	}

	defer file.Close()

}

// create 8 by 8 blocks

// shifted block minus 128 from every pixel value (lum) - center around 0

// discrete cosine transform 2 coefficients

// quantization table divide coefficient values by quantization value then round to nearest int

// huffman encoding

// serilize all 8 by 8 blocks into a line
