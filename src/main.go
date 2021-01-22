package main

import (
	"fmt"
)

func main() {
	var imageURL string = "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2965081049,3975994462&fm=26&gp=0.jpg"
	baiduImages(imageURL)
	fmt.Println("Hello, 世界")
}

func baiduImages(imageURL string) {
	var postData = [2] string {
		'from' => 'pc',
		'image' => $image,
	}
}
