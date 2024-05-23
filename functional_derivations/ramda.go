package main

import (
	"constraints"
	"fmt"
)

type Number interface{
	constraints.Integer|constraints.Float|constraints.Unsigned
}

func F(args... interface{})bool{
	return false
}

func T(args... interface{})bool{
	return true
}

const (
	__ = "@@functional/placeholder"
)

func isPlaceholder(x any)bool{
	if x == __ {
		return true
	}
	return false
}



func main(){
	fmt.Println(F(),T(),isPlaceholder(__))
}