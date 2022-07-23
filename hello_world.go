package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tshowimage [imgfile]")
		return
	}

	filename := os.Args[1]
	// window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)

	fmt.Println("Image: ", img.Total())
	// if img.Empty() {
	// 	fmt.Printf("Error reading image from: %v\n", filename)
	// 	return
	// }
	// for {
	// 	window.IMShow(img)
	// 	if window.WaitKey(1) >= 0 {
	// 		break
	// 	}
	// }

}
