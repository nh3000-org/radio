package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/rymdport/resize"
)

func main() {
	// open "test.jpg"
	file, err := os.Open("/home/oem/Videos/test.png")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(2000, 0, img, resize.Lanczos3)

	out, err := os.Create("/home/oem/Videos/test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
