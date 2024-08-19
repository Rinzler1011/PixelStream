package yuv

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
