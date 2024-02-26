package main

import (
	"fmt"
	"sort"
)

func main() {
	var ptr = []int{10,20,30,04,5,7,4,3,34}
	fmt.Println("the length of ptr",len(ptr))

	ptr = append(ptr , 1000000000,204,134)
	fmt.Println("after appending",ptr)

	ptr = append(ptr[2:6]) // 1st - skip , less than 3
	fmt.Println("after slicing",ptr)

	highscore := make([]int,5,5) // it means 
	fmt.Println("capacity before appending ", cap(highscore))
	highscore[0] = 1
	highscore[1] = 2
	highscore[2] = 4
	highscore[3] = 3
	highscore[4] = 5

	fmt.Println("before sorting highscore",highscore)
	highscore = append(highscore,6)
	sort.Ints(highscore)
	fmt.Println("ints are sorted",sort.IntsAreSorted(highscore))
	fmt.Println("highscore",highscore)
	//we can sort also using sort 
	fmt.Println("capacity after appeding", cap(highscore))

}