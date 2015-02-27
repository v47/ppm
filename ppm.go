//Package adds ppm image format support and provide method for recognise image formats.
package ppm

import (
	"bytes"
	_ "fmt"
	"image"
	"image/color"
	"io"
	"strconv"
	_ "time"
)

func PPMtoImage(input io.Reader) (Image image.Image, Comments string, err error) {

	TempBuff := new(bytes.Buffer)
	TempBuff.ReadFrom(input)
	buff := TempBuff.Bytes()

	var width int
	var height int

	position := 2

	if buff[position] == 0x0A && buff[position+1] == 0x23 {
		for i := position + 1; buff[i] != 0x0A; i++ {
			position = i
		}

		position += 1

		Comments = string(buff[3:position])
	}
	if buff[position] == 0x0A {
		var w []byte
		for i := position + 1; buff[i] != 0x20; i++ {
			w = append(w, buff[i])
			position = i
		}
		var h []byte
		for i := position + 2; buff[i] != 0x0A; i++ {
			h = append(h, buff[i])
			position = i
		}
		width, _ = strconv.Atoi(string(w))
		height, _ = strconv.Atoi(string(h))
	}
	buff = buff[position+6:]

	if width*height*3 == len(buff) {
		index := 0
		ImageC := image.NewRGBA(image.Rect(0, 0, width, height))
		for i := 0; i < height; i++ {
			for y := 0; y < width; y++ {
				ImageC.Set(y, i, color.RGBA{uint8(buff[index]), uint8(buff[index+1]), uint8(buff[index+2]), 255})
				index += 3
			}
		}
		Image = ImageC
	}

	return Image, Comments, nil
}
