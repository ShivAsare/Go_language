package main

import "fmt"

func main() {
	var Names [10]string
	Names[1] = "11"
	Names[2] = "12"
	Names[3] = "13"
	Names[4] = "14"
	Names[5] = "15"
	Names[6] = "16"
	Names[7] = "17"
	Names[8] = "18"
	Names[9] = "19"
	fmt.Println("Names =", Names)

	//Slice
	var students []string

	students = append(students, "vishal ")
	fmt.Println("students =", students[0])
}
