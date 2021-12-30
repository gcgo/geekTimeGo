package main

func main() {
	c := Cal{a: 1, b: 2, res: 0}
	c.Add()
	println(c.res)
}

type Cal struct {
	a, b int
	res  int
}

func (c *Cal) Add() {
	c.res = c.a + c.b
}
