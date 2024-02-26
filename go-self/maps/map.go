package main 

import "fmt"

func main() {
	maps := make(map[int]string)

	maps[0] = "q"
	maps[1] = "d"
	maps[2] = "p"
	maps[3] = "c"
	maps[4] = "b"
	maps[5] = "o"

	fmt.Println("the maps",maps)

	delete(maps,4)

	fmt.Println("after deleting 4",maps)

	// loops in map we use range 
	for key, value := range maps {
		fmt.Printf("for key %v, value is %v\n", key ,value) // basically %v takes value	
	}
}