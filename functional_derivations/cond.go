package main

import (
	"fmt"
)

type CondTU []struct {
	Comp func(string) bool
	Ap   func(string) int
}

func (c CondTU) Add(C func(string) bool, A func(string) int) CondTU {
	c = append(c, struct {
		Comp func(string) bool
		Ap   func(string) int
	}{
		Comp: C,
		Ap:   A,
	})
	return c
}

func (c CondTU) Call(inp string) int {
	for _, v := range c {
		if v.Comp(inp) {
			return v.Ap(inp)
		}
	}
	return 0
}

func cond(Comp []func(...string) bool, Ap []func(...string) int) func(...string) int {
	if len(Comp) != len(Ap) {
		panic("comp must be the same length as ap")
	}
	return func(inp ...string) int {
		for i, C := range Comp {
			if C(inp...) {
				return Ap[i](inp...)
			}
		}
		return 0
	}
}

type condTest func(string) int

func (c condTest) Add(C func(string) bool, A func(string) int) condTest {
	return func(inp string) int {
		if C(inp) {
			return A(inp)
		}
		return c(inp)
	}
}

func condMain() {
	strtoi := cond([]func(...string) bool{
		func(inp ...string) bool {
			return inp[0] == "0"
		},
		func(inp ...string) bool {
			return inp[0] == "1"
		},
	}, []func(...string) int{
		func(inp ...string) int {
			return 0
		},
		func(s ...string) int {
			return 1
		},
	})
	strtoi2 := CondTU{}.Add(func(t string) bool {
		return t == string("0")
	}, func(t string) int {
		return int(0)
	}).Add(func(t string) bool { return t == "1" }, func(t string) int { return 1 })

	strtoi3 := condTest(func(t string) int {
		return 0
	}).Add(func(t string) bool {
		return t == "1"
	}, func(t string) int {
		return 1
	})

	fmt.Printf("%+V x %+V x %+V\n", strtoi("0"), strtoi2.Call("0"), strtoi3("0"))
}
