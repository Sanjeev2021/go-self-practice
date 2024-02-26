package main

import "fmt"

func main() {

	var s string = "sanjeev"
	fmt.Println(s)
	fmt.Printf("%T\n" , s)

	var a bool = true
	fmt.Println(a)
	fmt.Printf("%T\n",a)

	var b float32 = 123.2
	fmt.Println(b)
	fmt.Printf("%T\n",b)

	var d int = 1
	fmt.Println(d)
	fmt.Printf("%T\n",d)

	var e = [2]int{10,20}
	fmt.Println(e)
	fmt.Printf("%T\n",e)

	var f = make([]int,3)
	fmt.Println(f)
	fmt.Printf("%t\n",f)

	fmt.Println("Enter your value")
	var l = 
	fmt.Scan(&l)
	fmt.Println("your value is", l)
	fmt.Printf("your value has type %T\n",l)
}

