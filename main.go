package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
)

func loadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

// Luminosity weights for 16-bit channel values, scaled by 2^16
// (0.299*65536, 0.587*65536, 0.114*65536).
const (
	wR = 19595
	wG = 38470
	wB = 7471
)

// avgPixel returns the average grayscale value (0-65535) of a block of pixels.
// The per-pixel luminosity weighting is folded into the block average, so the
// loop only accumulates raw channels and the weighted math happens once per
// block: avg(0.299R + 0.587G + 0.114B) == 0.299avg(R) + 0.587avg(G) + 0.114avg(B).
func avgPixel(img image.Image, x, y, w, h int) int {
	cnt, sumR, sumG, sumB := 0, 0, 0, 0
	max := img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			sumR += int(r)
			sumG += int(g)
			sumB += int(b)
			cnt++
		}
	}
	r, g, b := sumR/cnt, sumG/cnt, sumB/cnt
	return (r*wR + g*wG + b*wB) >> 16
}

// avgPixelColor returns the average grayscale value (0-65535) of a block of
// pixels, along with the average red, green and blue components (0-255).
func avgPixelColor(img image.Image, x, y, w, h int) (int, uint32, uint32, uint32) {
	cnt, sumR, sumG, sumB := 0, uint32(0), uint32(0), uint32(0)
	max := img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			sumR += r
			sumG += g
			sumB += b
			cnt++
		}
	}
	r, g, b := sumR/uint32(cnt), sumG/uint32(cnt), sumB/uint32(cnt)
	gray := int((0.299 * float64(r)) + (0.587 * float64(g)) + (0.114 * float64(b)))
	return gray, r >> 8, g >> 8, b >> 8
}

func main() {
	Args := os.Args[1:]

	useColor := false
	var pos []string
	for _, arg := range Args {
		switch arg {
		case "-c", "-color", "--color":
			useColor = true
		default:
			pos = append(pos, arg)
		}
	}

	if len(pos) > 0 && pos[0] == "help" {
		fmt.Println(`Usage: tiv [flags] <image_path> <scale_y> <scale_x>

Arguments:
  <image_path>   The path to the image file to be converted to ASCII art.
  <scale_y>      The vertical scale factor. Must be an integer.
  <scale_x>      The horizontal scale factor. Must be an integer.

Flags:
  -c, -color, --color
                 Render the ASCII art in color (uses 24-bit truecolor escape codes,
                 so it requires a terminal with truecolor support).

Description:
  This program converts an image to ASCII art. It requires three arguments:
  1. The path to the image file to be converted.
  2. The vertical scale factor (scale_y) which determines the number of image pixels to be averaged for each ASCII character in the vertical direction.
  3. The horizontal scale factor (scale_x) which determines the number of image pixels to be averaged for each ASCII character in the horizontal direction.`)
		os.Exit(0)
	}

	if len(pos) < 3 {
		fmt.Println("Missing Args, please type tiv help if you need to see usage")
		os.Exit(1)
	}

	scaleX, err := strconv.Atoi(pos[2])
	if err != nil {
		fmt.Println("X or Y scale not valid, please enter integers")
		os.Exit(1)
	}
	scaleY, err := strconv.Atoi(pos[1])
	if err != nil {
		fmt.Println("X or Y scale not valid, please enter integers")
		os.Exit(1)
	}

	img, err := loadImage(pos[0])
	if err != nil {
		panic(err)
	}
	ramp := `$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,"^'.`

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	max := img.Bounds().Max
	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			if useColor {
				c, r, g, b := avgPixelColor(img, x, y, scaleX, scaleY)
				fmt.Fprintf(out, "\x1b[38;2;%d;%d;%dm%c", r, g, b, ramp[len(ramp)*c/65536])
			} else {
				c := avgPixel(img, x, y, scaleX, scaleY)
				out.WriteByte(ramp[len(ramp)*c/65536])
			}
		}
		if useColor {
			out.WriteString("\x1b[0m")
		}
		out.WriteByte('\n')
	}
}
