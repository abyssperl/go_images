package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
  "math/rand"
	"os"
  "time"
)

func main() {
	width, height := 2048, 1024
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawGradient()

  // Draw a set of spirals randomly over the image
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
    x := float64(width) * rand.Float64()
    y := float64(height) * rand.Float64()
    color := color.RGBA{uint8(rand.Intn(255)),
				                uint8(rand.Intn(255)),
                        uint8(rand.Intn(255)),
                        255}
			                
		canvas.DrawSpiral(color, Vector{x, y})
  }

	outFilename := "spirals.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}

type Canvas struct {
	image.RGBA
}

func NewCanvas(r image.Rectangle) *Canvas {
	canvas := new(Canvas)
	canvas.RGBA = *image.NewRGBA(r)
	return canvas
}

func (c Canvas) DrawGradient() {
	size := c.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := color.RGBA{
				uint8(255 * x / size.X),
				uint8(255 * y / size.Y),
				55,
				255}
			c.Set(x, y, color)
		}
	}
}

func (c Canvas) DrawLine(color color.RGBA, from Vector, to Vector) {
	delta := to.Sub(from)
	length := delta.Length()
	x_step, y_step := delta.X/length, delta.Y/length
	limit := int(length + 0.5)
	for i := 0; i < limit; i++ {
		x := from.X + float64(i)*x_step
		y := from.Y + float64(i)*y_step
		c.Set(int(x), int(y), color)
	}
}

func (c Canvas) DrawSpiral(color color.RGBA, from Vector) {
	dir := Vector{0, 5}
	last := from
	for i := 0; i < 10000; i++ {
		next := last.Add(dir)
		c.DrawLine(color, last, next)
		dir.Rotate(0.03)
		dir.Scale(0.999)
		last = next
	}
}
