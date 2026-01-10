package domain

type CreativeFormat int

const (
	CreativeBanner CreativeFormat = iota
	CreativeVideo
	CreativeNative
)

type Creative struct {
	ID     int
	Format CreativeFormat
	Width  int
	Height int
	MIME   string
	Adm    string
}
