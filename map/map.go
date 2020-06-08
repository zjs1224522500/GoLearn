package main

import "fmt"

// Vertex Exported Struct
type Vertex struct {
	Lat, Long float64
}

func makeMap() {
	// The make function returns a map of the given type, initialized and ready for use.
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {
	// Map literals are like struct literals, but the keys are required.
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var mm = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(mm)
}

func mutateMaps() {
	m := make(map[string]int)

	// Insert
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// Update
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// Delete - After delete, the value is zero
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Get
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	makeMap()
	mapLiterals()
	mutateMaps()
}
