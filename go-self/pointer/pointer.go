package main

import "fmt"

func main() {

	a := "string"
	var ptr *string
	ptr = &a
	fmt.Println("the value inside pointer", ptr) // if it has no value is <nil> , if no value then nil 
	fmt.Println("what is the address of pointer", &ptr)
	fmt.Printf("the type of pointer %T\n",ptr) // type of pointer *int

	var pointer **string 
	pointer = &ptr
	fmt.Printf("THE TYPE OF POINTER %T\n",pointer)
	fmt.Println("the address of pointer", &pointer)
	fmt.Println("the value inside pointer", pointer) // this is giving runtime error because it has no value , so hence cant access it's value

	var c *string
	var d *string
	d = &a
	c = &(*d)
	fmt.Printf("the type of pointer %T\n",c)
	fmt.Println("the value inside c ", c) // this will only give me address of d 
	fmt.Println("the address of c", &c) // this will give me address of c
	fmt.Println("the address of d", &d)
	fmt.Println("the value in c", *c)
}