package main

import "fmt"

type Tuple struct {
	first, second int
}

func main() {
	simpleTuple()
	mapLiteral()
	cleanMapLiteralSyntax()
}

func simpleTuple() {
	var m = make(map[string]Tuple)
	m["First"] = Tuple{10, 20}
	m["Second"] = Tuple{40, 80}
	fmt.Println(m)
}

func mapLiteral() {
	var m = map[string]Tuple{
		"First": Tuple{
			10, 20,
		},
		"Second": Tuple{
			40, 80,
		},
	}
	fmt.Println(m["Second"])
}

//Map values struct type is omited
func cleanMapLiteralSyntax() {
	var m = map[string]Tuple{
		"First": {10, 20},
		"Second": {40, 80},
	}
	fmt.Println("Berfor update", m)
	
	m["First"] = Tuple{20, 40}
	fmt.Println("After update", m)

	delete(m, "First")
	fmt.Println("After delete", m)

	//Test if map key exists
	key := "Second"
	_, exists := m[key]
	
	if !exists {
		fmt.Println(key, "does not exist")
	} else {
		fmt.Println(key, "exists")
	}
}