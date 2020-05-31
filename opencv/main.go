/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-04-20 14:21 
# @File : main.go
# @Description : 
# @Attention : 
*/

package main

import (
	"gocv.io/x/gocv"
)
//

/**

docker run -v "$GOPATH":/Users/joker/go --rm -v "$PWD":/Users/joker/go/src/examples/opencv -w /Users/joker/go/src/examples/opencv  -e GOOS="linux" -e GOARCH="amd64" 3f779b8e2183 go build -v
 */

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
