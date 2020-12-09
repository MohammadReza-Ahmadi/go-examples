package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(power(3,2,10),power(3,3,20))

}

func power(x,n,lim float64) float64{
	if v:=math.Pow(x,n); v < lim {
		return v
	}
	return lim
	//execution of this line shows: undefined: v
	//return v
}
