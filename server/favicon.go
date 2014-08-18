package server

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"strconv"
)

var (
	blue   = color.RGBA{0x15, 0x8C, 0xBA, 0xFF}
	icon16 = iconBytes(16, blue)
)

func favicon(w http.ResponseWriter, r *http.Request) {
	s, err := strconv.Atoi(r.URL.Query().Get("s"))

	// Return default icon if error,
	// less than 16 px or above 1024 px
	if err != nil || s <= 16 || s > 1024 {
		w.Write(icon16)
		return
	}

	w.Write(iconBytes(s, blue))
}

func fill(i *image.RGBA, c color.RGBA) *image.RGBA {
	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)

	return i
}

func iconBytes(s int, c color.RGBA) []byte {
	return pngBytes(iconImage(s, c))
}

func iconImage(s int, c color.RGBA) *image.RGBA {
	return fill(image.NewRGBA(image.Rect(0, 0, s, s)), blue)
}

func pngBytes(i *image.RGBA) []byte {
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)

	png.Encode(w, i)

	w.Flush()

	return b.Bytes()
}
