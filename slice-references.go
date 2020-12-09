package main

import "fmt"

func main() {
	names:=[4]string{
		"Apple",
		"Google",
		"Oracle",
		"Netflix"}

	fmt.Println(names)
	a:=names[0:2]
	b:=names[1:3]
	fmt.Println("a=",a,", b=",b)

	b[0]="XXX"
	fmt.Println("a=",a,", b=",b)
	fmt.Println("names=",names)

}
