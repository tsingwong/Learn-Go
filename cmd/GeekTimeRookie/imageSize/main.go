/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-10-29 11:55:19
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-10-29 12:00:11
 */
package main

import (
	"fmt"
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	width, height, err := imageSize(os.Args[1])
	if err != nil {
		fmt.Println("get image size error: ", err)
		return
	}
	fmt.Printf("image size: [%d, %d]\n", width, height)
}

func imageSize(imageFile string) (int, int, error) {
	f, _ := os.Open(imageFile)
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return 0, 0, err
	}

	b := img.Bounds()
	return b.Max.X, b.Max.Y, nil
}
