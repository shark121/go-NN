package helpers

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	mat "gonum.org/v1/gonum/mat"
)

func DrawImage(width int, height int, pixels mat.Dense) {

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.Gray{Y: uint8(pixels.At(y, x))})
		}
	}

	file, err := os.Create("output.jpg")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close() // Make sure to close the file when we're done.

	// 4. Encode the image as a JPEG and write it to the file.
	// The third argument is the quality option (0-100). nil gives a default.
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatalf("Failed to encode JPEG: %v", err)
	}

	log.Println("Successfully created output.jpg")
}
