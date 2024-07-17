package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
)


func loadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil{
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func grayScale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int((0.299 * float64(r)) + (0.587 * float64(g)) + (0.114 * float64(b)) )
}

func avgPixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += grayScale(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}
func main(){
	Args := os.Args


	if Args[1]== "help"{
		fmt.Println(`Usage: tiv <image_path> <scale_y> <scale_x>

Arguments:
  <image_path>   The path to the image file to be converted to ASCII art.
  <scale_y>      The vertical scale factor. Must be an integer.
  <scale_x>      The horizontal scale factor. Must be an integer.

Description:
  This program converts an image to ASCII art. It requires three arguments:
  1. The path to the image file to be converted.
  2. The vertical scale factor (scale_y) which determines the number of image pixels to be averaged for each ASCII character in the vertical direction.
  3. The horizontal scale factor (scale_x) which determines the number of image pixels to be averaged for each ASCII character in the horizontal direction.`)
	}

	if len(Args) < 3 {
		fmt.Println("Missing Args, please type tiv help if you need to see usage")
		os.Exit(1)
	}



	scaleX, err := strconv.Atoi(Args[3])
	scaleY, err := strconv.Atoi(Args[2])

	if err !=nil{
		fmt.Println("X or Y scale not valid, please enter integers")
		os.Exit(1)
	}

	
	img, err := loadImage(Args[1])
	if err != nil{
		panic(err)
	}
	ramp := `$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,"^'.`

	max := img.Bounds().Max
	for y := 0; y<max.Y; y+= scaleX{
		for x := 0; x<max.X; x +=scaleY{
			c:=avgPixel(img, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[len(ramp) * c/65536]))
		}
		fmt.Println()
	}
}
