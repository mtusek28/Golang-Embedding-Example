package Shapes

import "math"

// Shape --> General shape with generic properties used as a base for shape inheritance
type Shape struct {
	Area, Perimiter float64
	Shapetype       string
}

// Square --> A class of a square that has the extension of a shape and Rectangle but with new properties such as sidelength
type Square struct {
	Shape
	Rectangle
	Sidelength float64
}

// Rectangle --> A class for a square that has the extension of a shape but with new properties such as Height, Width
type Rectangle struct {
	Shape
	Height, Width float64
}

// Triangle --> A class for a triangle that is the extension of a shape but with the new properties of a Base, Height, and the Hypotenuse
type Triangle struct {
	Shape
	Height, Base, Hypotenuse float64
	isRightTriangle          bool
}

// Circle --> A class for a circle that extends upon the class Shape but has the new properties of a Radius and a Diameter
type Circle struct {
	Shape
	Radius, Diameter, Circumference float64
}

// SetObjectDataSquare --> Sets the object data for this specific object
func SetObjectDataSquare(s *Square, Sidelength float64) {
	s.Sidelength = Sidelength
	s.Shapetype = "Square"
	s.Area = math.Pow(Sidelength, 2)
	s.Perimiter = Sidelength * 4
}

// SetObjectDataCircle --> Sets the object data for this specific object
func SetObjectDataCircle(c *Circle, Radius float64) {
	c.Radius = Radius
	c.Diameter = 2 * Radius
	c.Shapetype = "Circle"
	c.Area = math.Pi * math.Pow(Radius, 2)
	c.Circumference = 2 * math.Pi * Radius
}

// SetObjectDataRectangle --> Sets the object data for this specific object
func SetObjectDataRectangle(r *Rectangle, Height float64, Width float64) {
	r.Height = Height
	r.Width = Width
	r.Shapetype = "Rectangle"
	r.Area = Height * Width
	r.Perimiter = 2*Height + 2*Width
}

// SetObjectDataTriangle --> Sets the object data for this specific object
func SetObjectDataTriangle(t *Triangle, Height float64, Base float64, isRight bool) {
	t.Height = Height
	t.Base = Base
	t.Shapetype = "Triangle"
	t.isRightTriangle = isRight
	t.Area = (Height * Base) / 2

	if isRight {
		t.Hypotenuse = math.Sqrt((math.Pow(t.Height, 2)) + (math.Pow(t.Base, 2)))
		t.Perimiter = t.Height + t.Base + t.Hypotenuse
	} else {
		t.Hypotenuse = 0
	}
}
