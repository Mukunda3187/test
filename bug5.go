package main

import "fmt"

func classify(x int) string {
	if x > 0 {
		return "positive"
	} else {
		return "negative"
	}
	fmt.Println("dead code, but Go is excluded from this check")
}

func main() {
	fmt.Println(classify(5))
}
