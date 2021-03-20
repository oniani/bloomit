package main

import "fmt"

// Test
func test() {
	bf := New(5, 0.25)
	bf.Add("sada")
	res := bf.Check("ssada")
	fmt.Println(res)
}

func main() {
	// test()
}
