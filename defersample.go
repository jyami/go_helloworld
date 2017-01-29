package main

import (
	"fmt"
	"os"
)

func mainDeferSample() {
	file, err := os.Open("hoge.txt")
	if err != nil {
		fmt.Println("File open error: ", err)
		return
	}
	defer file.Close()

	fmt.Println("some process")
}
