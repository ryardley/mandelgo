package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	start := time.Now()
	generateMandelBrot()
	elapsed := time.Since(start)
	log.Printf("Generation took %s", elapsed)

	port := "8080"
	directory, err := os.Getwd()

	if err != nil {
		return
	}

	http.Handle("/", http.FileServer(http.Dir(directory)))
	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func belongs(x float64, y float64) float64 {
	realComponentOfResult := x
	imaginaryComponentOfResult := y
	maxIterations := 100
	for i := 0; i < maxIterations; i++ {
		tempRealComponent := realComponentOfResult*realComponentOfResult - imaginaryComponentOfResult*imaginaryComponentOfResult + x
		tempImaginaryComponent := 2*realComponentOfResult*imaginaryComponentOfResult + y
		realComponentOfResult = tempRealComponent
		imaginaryComponentOfResult = tempImaginaryComponent

		// Return a number as a percentage
		if math.Sqrt(realComponentOfResult*realComponentOfResult+imaginaryComponentOfResult*imaginaryComponentOfResult) > 2 {
			return (float64(i)/float64(maxIterations))*99.0 + 1.0
		}
	}
	return 0.0
}

func generateMandelBrot() {
	Width := 1400
	Height := 700
	ScaleW := 6.0
	ScaleH := 3.0

	// background := color.RGBA{0, 0xFF, 0, 0xCC}
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))

	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			x := ScaleW*float64(i)/float64(Width) - ScaleW/2.0
			y := ScaleH*float64(j)/float64(Height) - ScaleH/2.0

			val := belongs(x, y)

			if val == 0.0 {
				img.Set(i, j, color.RGBA{0x00, 0x00, 0x00, 0xff})
			} else {
				img.Set(i, j, colorful.Hsv(60.0, 1.0, (val/100.0)))
			}
		}
	}

	outputFile, err := os.Create("images/test.png")
	if err != nil {
		fmt.Println("The was an error")
	}

	png.Encode(outputFile, img)
	outputFile.Close()
}
