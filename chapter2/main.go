package main

import "fmt"



func smallest() {

	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}

	i := 0
	smallest := x[0]
	for i < len(x) {
		if x[i] < smallest {
			smallest = x[i]
		}
		i += 1

	}

	fmt.Println(smallest)	
}


func main() {

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	// first we try to get the value from the map. Then, if its successful
	// we run the code inside of the block

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	elements2 := map[string]string{

		"H": "Hydrogen",
		"He": "Helium",
	}

	fmt.Println(elements2)


	smallest()
}


