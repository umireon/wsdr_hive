package controllers

import (
	"bytes"
	"crypto/rc4"
	"fmt"
	"github.com/revel/revel"
	"github.com/umireon/wsdr_hive/app"
	_ "golang.org/x/image/bmp"
	"image"
	"os"
)

type Event struct {
	*revel.Controller
}

func (c Event) Activate() revel.Result {
	numListener := 0
	for {
		select {
		case app.Event.Activate <- struct{}{}:
			numListener += 1
		default:
			type message struct {
				NumListener int
			}
			return c.RenderJson(message{numListener})
		}
	}
}

func (c Event) ListenActivate() revel.Result {
	<-app.Event.Activate
	return c.RenderJson(true)
}

func (c Event) Frame_tx() revel.Result {
	buf := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(buf)
	fmt.Println(c.Request.ContentType)
	fmt.Println(buf)
	return c.RenderJson(len(buf))
}

func (c Event) DataFetch() revel.Result {
	r, err := os.Open("public/img/splash.bmp")
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
			buf.Write([]byte{byte(r), byte(g), byte(b)})
		}
	}
	cip, _ := rc4.NewCipher([]byte{42})
	bbuf := buf.Bytes()
	cip.XORKeyStream(bbuf, bbuf)
	return c.RenderText(string(bbuf))
}
