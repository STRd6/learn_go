package main

import "fmt"

func main() {
	var foo = "hello world"
	const pi = 3.14159
	bar := "yolo"
	baz := true

	fmt.Printf("%v\n%v\n%v\n%v\n",
		foo,
		pi,
		bar,
		baz,
	)
}
