package util

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
	"golang.org/x/crypto/ssh/terminal"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

var pixels = []byte(".,-~:;=!*#@$")

func Convert(filePath string) string {
	img, err := readFile(filePath)
	if err != nil {
		log.Fatal("Read file error: " + err.Error())
	}
	img = resizeImg(img)
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
	brightness := (pixel.R + pixel.G + pixel.B)
	scale := float64(255 * 3 / (len(pixels) - 2))

	index := int(math.Round(float64(brightness) / scale))
	return string(pixels[index])
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func resizeImg(img image.Image) image.Image {
	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())
	width, height, _ := terminal.GetSize(0)
	ratio := float64(height) / imgHeight
	if (imgWidth * ratio / 0.5) < float64(width) {
		return resize.Resize(uint(imgWidth*ratio/0.5), uint(imgHeight*ratio), img, resize.Lanczos3)
	}
	ratio = float64(width) / imgWidth
	return resize.Resize(uint(imgWidth*ratio), uint(imgHeight*ratio), img, resize.Lanczos3)
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
