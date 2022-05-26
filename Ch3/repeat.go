package main

import "fmt"

const repeatCount = 5

func Repeat(letter string) string {

	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += letter
	}

	return repeated

}

func main() {
	fmt.Println(Repeat("a"))
}
