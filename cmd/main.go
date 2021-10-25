package main

import (
	"fmt"

	"github.com/vSterlin/text-color/pkg/textcolor"
)

func main() {
	abc := "abc"
	abc, _ = textcolor.AddColor(abc, textcolor.RED)

	fmt.Println(abc)
	fmt.Println("def")

}
