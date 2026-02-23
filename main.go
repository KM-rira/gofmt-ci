package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}

	// -s(simplify) で簡略化される（0:len(arr) → :）
	s := arr[0:len(arr)]
	fmt.Println(s)

	// -s で簡略化される（int(1) → 1 など）
	x := []int{int(1), int(2), int(3)}
	fmt.Println(x)

	// これは通常の gofmt でも整形される（スペースなど）
	y  :=  1
	fmt.Println(y)
}
