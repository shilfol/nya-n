package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nyan := []string{"に", "ゃ", "ー", "ん", ""}
	rand.Seed(time.Now().UnixNano())

	nyaon := 0
	for {
		if callNyan(nyan) {
			break
		}
		nyaon++
	}
	fmt.Println("nyannyan: ", nyaon)
}

func callNyan(nyan []string) bool {
	nyaa := "にゃーん"
	unya := ""
	for i := 0; i < 4; i++ {
		unya += nyan[rand.Intn(len(nyan))]
	}
	fmt.Println(unya)

	if unya == nyaa {
		return true
	} else {
		return false
	}
}
