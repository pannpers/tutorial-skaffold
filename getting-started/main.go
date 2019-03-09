package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello Kaji!")

		time.Sleep(time.Second * 1)
	}
}
