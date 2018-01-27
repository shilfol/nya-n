package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nyan := []string{"に", "ゃ", "ー", "ん", ""}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 22; i++ {
		callNyan(nyan)
	}
}

func callNyan(nyan []string) {
	for i := 0; i < 4; i++ {
		fmt.Print(nyan[rand.Intn(len(nyan))])
	}
	fmt.Println()

}
