package main

import "fmt"

func main() {

	i,j:=42,2701
	p:=&i

	//dereferencing
	fmt.Println("i pointer value=",*p)
	//indirecting
    *p=21
    fmt.Println("*p value after set=",*p)
    fmt.Println("i value after set=",i)

    p=&j
    *p=(*p / 37)+1
    fmt.Println("j value after set=",j)

	
}
