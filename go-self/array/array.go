package main

import "fmt"

func main() {
// there are total 5 ways to make array - 
	// var arr [3]int
	// arr1 := [2]int{}
	// fmt.Println("the value in array", arr1)
	// fmt.Printf("the type of array %T\n",arr1)
	// fmt.Println("the address of array", &arr1)
	// fmt.Println("the length of array",len(arr1))


	// if i dont want array lenght to be inferred , oh sexy 
	arr := [...]int{1,2,3}
	 fmt.Println("the value in array", arr)
	 fmt.Printf("the type of array %T\n",arr)
	 fmt.Println("the address of array", &arr)
	 fmt.Println("the length of array",len(arr))

	// if i want to use pointer to array
	arr5:= new([5]int)
	fmt.Println("the pointer to arr5 is ", *arr5)
	fmt.Printf("the type of arry %T\n", arr5) // *[5]int
	fmt.Println("the address of arr5", &arr5)
	fmt.Println("the length of arr5", len(arr5))

	arr8 := [3]int{1,2,3}

	var ptr1 (*[3]int)
	ptr1 = &arr8
	fmt.Println("the value in ptr1" , *ptr1)
	fmt.Println("the address of ptr" , &ptr1)
	fmt.Printf("the type of pointer %T\n",ptr1 )

	var ptr2 (**[3]int)
	ptr2 = &ptr1
	fmt.Println("the value of ptr2", *ptr2)
	fmt.Println("for my understanding", ptr2)
	fmt.Println("the address of ptr2", &ptr2)

	var string = [3]string{"teri","jai","ho"}
	// strings :=[3]string{"yo"}
    // string = append(string , "lol")
	fmt.Println("the value of teri ma", string)
}