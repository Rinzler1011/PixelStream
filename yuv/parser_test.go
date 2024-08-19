package yuv

import (
	"log"
	"os"
	"testing"
)

func TestYUV(t *testing.T) {

	width := 640
	height := 480
	ySize := width * height
	uSize := (width / 2) * (height / 2)
	vSize := (width / 2) * (height / 2)

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

	for i := range yuv.Y {
		yuv.Y[i] = byte(i % 256)
	}

	for i := range yuv.U {
		yuv.U[i] = byte(i % 256)
	}

	for i := range yuv.V {
		yuv.V[i] = byte(i % 256)
	}

	f, err := os.Create("testing/test.yuv")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(yuv.Y)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(yuv.U)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(yuv.V)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}
