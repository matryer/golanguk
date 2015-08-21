package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(3) + 1)
}
