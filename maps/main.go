package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["xyz"] = 10
	m["pqr"] = 100

	val, ok := m["pqr"]
	fmt.Println("val =", val, "ok =", ok)

	val, ok = m["xyz"]
	if ok {
		fmt.Println("val =", val, "ok =", ok)
	} else {
		fmt.Println("jadhfijvhijuhi")
	}

}
