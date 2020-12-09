package main

import "fmt"

func main() {
	fmt.Println(Vertex{10,20})
	v:= Vertex{5,10}
	fmt.Println("v=",v)
	v.X = 41
	fmt.Println("v.X=",v.X,", v.Y=",v.Y)
	p:= &v
	p.X = 1e9
	fmt.Println("p.X value=",(*p).X)
	fmt.Println("v.X=",v.X,", v.Y=",v.Y)
}


type Vertex struct {
	X int
	Y int
}