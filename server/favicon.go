package server

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
)

var (
	blue = color.RGBA{0x15, 0x8C, 0xBA, 0xFF}
	icon = pngBytes(fill(image.NewRGBA(image.Rect(0, 0, 16, 16)), blue))
)

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Write(icon)
}

func fill(i *image.RGBA, c color.RGBA) *image.RGBA {
	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)

	return i
}

func pngBytes(i *image.RGBA) []byte {
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)

	png.Encode(w, i)

	w.Flush()

	return b.Bytes()
}
