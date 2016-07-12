package main

import "fmt"

// Move takes number of disks and start, to, via polls and moves disks
func Move(n, from, to, via int) {
	if n <= 0 {
		return
	}

	Move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	Move(n-1, via, to, from)
}

// Hanoi takes number of disks and calls Move func
func Hanoi(n int) {
	fmt.Println("Number of disks:", n)
	Move(n, 1, 2, 3)
}

func main() {
	Hanoi(3)
}
