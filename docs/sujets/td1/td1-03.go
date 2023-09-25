package main

import (
	"math"
)

type Point2D struct {
	x, y float64
}

func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{x, y}
}

func (p *Point2D) X() float64 {
	return p.x
}

func (p *Point2D) Y() float64 {
	return p.y
}

func (p *Point2D) SetX(x float64) {
	p.x = x
}

func (p *Point2D) SetY(y float64) {
	p.y = y
}

func (p *Point2D) Clone() *Point2D {
	return &Point2D{p.x, p.y}
}

func (p *Point2D) Module() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type Rectangle struct {
	topLeft, bottomRight Point2D
}

func NewRectangle(topLeft, bottomRight Point2D) *Rectangle {
	return &Rectangle{topLeft, bottomRight}
}

// func NewRectangle(topLeft, bottomRight *Point2D) *Rectangle {
// 	return &Rectangle{*topLeft.Clone(), *bottomRight.Clone()}
// }

func (r *Rectangle) TopLeft() *Point2D {
	return &r.topLeft
}

func (r *Rectangle) BottomRight() *Point2D {
	return &r.bottomRight
}

func (r *Rectangle) SetTopLeft(p Point2D) {
	r.topLeft = p
}

func (r *Rectangle) SetBottomRight(p Point2D) {
	r.bottomRight = p
}

type Sprite struct {
	position   Point2D
	hitbox     Rectangle
	zoom       float64
	bitmapFile string
}

func NewSprite(position Point2D, hitbox Rectangle, zoom float64, bitmapFile string) *Sprite {
	return &Sprite{position, hitbox, zoom, bitmapFile}
}

func (s *Sprite) Position() *Point2D {
	return &s.position
}

func (s *Sprite) Hitbox() *Rectangle {
	return &s.hitbox
}

func (s *Sprite) Zoom() float64 {
	return s.zoom
}

func (s *Sprite) BitmapFile() string {
	return s.bitmapFile
}

func (s *Sprite) Move(p Point2D) {
	s.position = p
}

func (s *Sprite) Collision(other *Sprite) Rectangle {
	// Calcul du rectangle de collision (ceci est une simplification, d'autres logiques pourraient être nécessaires)
	left := math.Max(s.hitbox.topLeft.X(), other.hitbox.topLeft.X())
	top := math.Min(s.hitbox.topLeft.Y(), other.hitbox.topLeft.Y())
	right := math.Min(s.hitbox.bottomRight.X(), other.hitbox.bottomRight.X())
	bottom := math.Max(s.hitbox.bottomRight.Y(), other.hitbox.bottomRight.Y())

	if left > right || top < bottom {
		return Rectangle{} // Pas de collision, retourne un rectangle vide
	}
	return *NewRectangle(*NewPoint2D(left, top), *NewPoint2D(right, bottom))
}

// func main() {
// 	// Vous pouvez tester vos structures et méthodes ici.
// }
