package main

import "math"

//その形の面積を調べる関数が実装されている
type Shape interface {
	Area() float64
}

//長方形の情報
type Rectangle struct {
	Width  float64
	Height float64
}

//長方形の面積を返す関数
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//円の情報
type Circle struct {
	Radius float64
}

//円の面積を返す関数
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Height * t.Width / 2
}

//長方形の外周
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

////長方形の面積
//func Area(rectangle Rectangle) float64 {
//	return rectangle.Width * rectangle.Height
//}
