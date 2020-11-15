package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(89)
	for _, r := range p[:5] {
		fmt.Println(r + 1)
	}
}
