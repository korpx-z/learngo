//create two structs that find their area via functions made
//create interface called shape that enables any using the functions made to inherit type shape
//use printArea against both types to prove interface inheritance.

package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	var t triangle
	t.height = 2.02
	t.base = 2.32

	var s square
	s.sideLength = 4.04

	printArea(t)
	printArea(s)

}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}
func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}
func printArea(sh shape) {
	fmt.Printf("%v", sh.getArea())
}
