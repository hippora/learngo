package main

import (
	"fmt"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount1 returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {
	var count int
	var i uint
	for i = 0; i < 63; i++ {
		count += int(byte((x >> i) & 1))
	}
	return count
}

func main() {
	start := time.Now()
	fmt.Println(PopCount(9876543))
	fmt.Println(time.Since(start).Nanoseconds())
	start = time.Now()
	fmt.Println(PopCount1(9876543))
	fmt.Println(time.Since(start).Nanoseconds())

}
