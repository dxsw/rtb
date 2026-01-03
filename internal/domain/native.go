package domain

type Native struct {
	Version string

	Assets []NativeAsset
}

type NativeAsset struct {
	ID       int
	Required bool

	Title *NativeTitle
	Image *NativeImage
	Data  *NativeData
}

type NativeTitle struct {
	Length int
}

type NativeImage struct {
	Type   int
	Width  int
	Height int
}

type NativeData struct {
	Type   int
	Length int
}
