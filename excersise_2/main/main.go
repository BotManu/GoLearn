package main

import "fmt"

const (
	_    = iota
	base = 10 + iota
	one
	two
	three
	four
	five
)

type ByteSize int64

const (
	_           = iota //ignore first value by assigning it a blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Println(one, two, three, four, five)
	fmt.Printf("%d \t \t %b\n", KB, KB)
	fmt.Printf("%d \t \t %b\n", MB, MB)
	fmt.Printf("%d \t \t %b\n", GB, GB)
	fmt.Printf("%d \t \t %b\n", TB, TB)
	fmt.Printf("%d \t \t %b\n", PB, PB)
}
