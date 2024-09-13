package yuv

import (
	"log"
	"os"
	"testing"
)

func TestYUV(t *testing.T) {

	pixels, err := ReadRGBFile("testing/testimage.rgb")
	if err != nil {
		log.Fatal(err)
	}

	yuv, err := RGBToYUV(pixels, 1920, 1080, 444)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("testing/test.yuv")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write(yuv.Y); err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(yuv.U); err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(yuv.V); err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}
