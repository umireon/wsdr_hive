package controllers

import (
	"bytes"
	"crypto/rc4"
	"github.com/revel/revel"
	_ "golang.org/x/image/bmp"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
)

type Data struct {
	*revel.Controller
}

type Metadata struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func openFile(id int) (*os.File, error) {
	var file string
	switch id {
	case 0:
		file = "public/img/splash.bmp"
	case 1:
		file = "public/img/glyphicons-halflings.png"
	case 2:
		file = "public/img/glyphicons-halflings-white.png"
	case 3:
		file = "public/img/favicon.png"
	}
	return os.Open(file)
}

func (c Data) Meta() revel.Result {
	id, _ := strconv.Atoi(c.Params.Values["id"][0])
	r, err := openFile(id)
	if err != nil {
		panic(err)
	}
	conf, _, err := image.DecodeConfig(r)
	if err != nil {
		panic(err)
	}
	return c.RenderJson(Metadata{
		Width:  conf.Width,
		Height: conf.Height,
	})
}

func (c Data) Fetch() revel.Result {
	id, _ := strconv.Atoi(c.Params.Values["id"][0])
	r, err := openFile(id)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(r)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	buf := new(bytes.Buffer)
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 1 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 1 {
			r, g, b, _ := img.At(x, y).RGBA()
			buf.Write([]byte{byte(r >> 8), byte(g >> 8), byte(b >> 8)})
		}
	}
	cip, _ := rc4.NewCipher([]byte{42})
	bbuf := buf.Bytes()
	cip.XORKeyStream(bbuf, bbuf)
	return c.RenderText(string(bbuf))
}
