package main

import "./Shapes"
import "fmt"

func main() {
	square := new(Shapes.Square)
	rectangle := new(Shapes.Rectangle)
	circle := new(Shapes.Circle)
	triangle := new(Shapes.Triangle)

	Shapes.SetObjectDataSquare(square, 10)
	Shapes.SetObjectDataRectangle(rectangle, 10, 10)
	Shapes.SetObjectDataCircle(circle, 10)
	Shapes.SetObjectDataTriangle(triangle, 10, 10, true)

	fmt.Printf("Shape Type: %v; Area: %v; Sidelength: %v; Perimiter: %v;\n", square.Shapetype, square.Area, square.Sidelength, square.Perimiter)
	fmt.Printf("Shape Type: %v; Area: %v; Height: %v; Width: %v; Perimiter: %v;\n", rectangle.Shapetype, rectangle.Area, rectangle.Height, rectangle.Width, rectangle.Perimiter)
	fmt.Printf("Shape Type: %v; Area: %v; Radius: %v; Diameter: %v; Circumference: %v;\n", circle.Shapetype, circle.Area, circle.Radius, circle.Diameter, circle.Circumference)
	fmt.Printf("Shape Type: %v; Area: %v; Height: %v; Base: %v; Hypotenuse: %v;\n", triangle.Shapetype, triangle.Area, triangle.Height, triangle.Base, triangle.Hypotenuse)

}
