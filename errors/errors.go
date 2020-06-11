package main

import (
	"fmt"
	"math"
	"time"
)

// MyError struct
type MyError struct {
	When time.Time
	What string
}

// The error type is a built-in interface similar to fmt.Stringer:
// type error interface {
//     Error() string
// }
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func testRunError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// Small Diffeneces
const Small = 1.0 / (1 << 20)

// ErrNegativeSqrt Error
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

// Sqrt Exported method
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		diff := (z*z - x) / (2 * z)
		if math.Abs(diff) < Small {
			break
		}
		z -= diff

		fmt.Println(z)
	}
	return z, nil
}

func errorExercise() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func main() {
	testRunError()
	errorExercise()
}
