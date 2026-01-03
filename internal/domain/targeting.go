package domain

type Bitmap []uint64

type Targeting struct {
	GeoBitmap    Bitmap
	DeviceBitmap Bitmap
	SiteBitmap   Bitmap
	TimeBitmap   Bitmap
}
