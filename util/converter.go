package util

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

var pixels = []byte("`^,:;Il!i~+_-?)(|tfjrxnuvczUJCmwqpao*#W&8%B@$")

func Convert(filePath string) string {
	img, err := readFile(filePath)
	if err != nil {
		log.Fatal("Read file error: " + err.Error())
	}
	img = resizeImg(img, 100, 30)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	res := ""
	for i := 0; i < int(height); i++ {
		raw := ""
		for j := 0; j < int(width); j++ {
			pixel := rgbaToPixel(img.At(j, i).RGBA())
			raw += convertPixelToChar(&pixel)
		}
		res += raw + "\n"
	}

	return res
}

func convertPixelToChar(pixel *Pixel) string {
	brightness := (pixel.R + pixel.G + pixel.B) * pixel.A / 255
	scale := float64(255 * 3 / (len(pixels) - 1))
	index := int(math.Round(float64(brightness) / scale))
	return string(pixels[index])
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func resizeImg(img image.Image, w int, h int) image.Image {
	return resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
}

func readFile(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(f)
	return image, err
}
