package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// import "fmt"
// import "math"

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	rand.Seed(time.Now().Unix())
	fmt.Println("(rand.Seed) My favorite number is", rand.Intn(10))
	fmt.Println("math.pi", math.Pi)
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
