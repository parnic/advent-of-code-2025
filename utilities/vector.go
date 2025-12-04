package utilities

import "math"

type Vec2[T Number] struct {
	X T
	Y T
}

type Vec3[T Number] struct {
	X T
	Y T
	Z T
}

type Vec2i Vec2[int]

var (
	ZeroVec2  = Vec2i{}
	Left      = Vec2i{X: -1, Y: 0}
	Right     = Vec2i{X: 1, Y: 0}
	Up        = Vec2i{X: 0, Y: -1}
	Down      = Vec2i{X: 0, Y: 1}
	FourWay   = []Vec2i{Right, Left, Up, Down}
	EightWay  = []Vec2i{Up, Up.AddVec(Right), Right, Right.AddVec(Down), Down, Down.AddVec(Left), Left, Left.AddVec(Up)}
	Diagonals = []Vec2i{Up.AddVec(Right), Down.AddVec(Right), Up.AddVec(Left), Down.AddVec(Left)}
)

func (v Vec2[T]) Dot(other Vec2[T]) T {
	return (v.X * other.X) + (v.Y * other.Y)
}

func (v Vec2[T]) Len() T {
	return T(math.Sqrt(float64(v.LenSquared())))
}

func (v Vec2[T]) LenSquared() T {
	return (v.X * v.X) + (v.Y * v.Y)
}

func (v Vec2[T]) To(other Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec2i) AddVec(other Vec2i) Vec2i {
	return Vec2i{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec2[T]) AngleBetween(other Vec2[T]) float64 {
	rad := math.Atan2(float64(other.Y-v.Y), float64(other.X-v.X))
	return rad * 180 / math.Pi
}

func (v Vec2[T]) Equals(other Vec2[T]) bool {
	return v.X == other.X &&
		v.Y == other.Y
}

func (v Vec2[T]) ManhattanDistance(other Vec2[T]) T {
	return T(math.Abs(float64(v.X-other.X)) + math.Abs(float64(v.Y-other.Y)))
}

func (v Vec2i) IsWithinRange(min Vec2i, max Vec2i) bool {
	return v.X >= min.X && v.X <= max.X && v.Y >= min.Y && v.Y <= max.Y
}

func (v Vec2i) GetBoundedNeighbors(min Vec2i, max Vec2i) []Vec2i {
	ret := make([]Vec2i, 0, len(EightWay))
	for _, n := range EightWay {
		var pt = v.AddVec(n)
		if !pt.IsWithinRange(min, max) {
			continue
		}
		ret = append(ret, pt)
	}
	return ret
}

func (v Vec2i) GetBoundedOrthogonalNeighbors(min Vec2i, max Vec2i) []Vec2i {
	ret := make([]Vec2i, 0, len(FourWay))
	for _, n := range FourWay {
		var pt = v.AddVec(n)
		if !pt.IsWithinRange(min, max) {
			continue
		}
		ret = append(ret, pt)
	}
	return ret
}

func VecBetween[T Number](a, b Vec2[T]) Vec2[T] {
	return a.To(b)
}

func ManhattanDistance[T Number](a, b Vec2[T]) T {
	return a.ManhattanDistance(b)
}

func (v Vec3[T]) Dot(other Vec3[T]) T {
	return (v.X * other.X) + (v.Y * other.Y) + (v.Z * other.Z)
}

func (v Vec3[T]) Len() T {
	return T(math.Sqrt(float64(v.LenSquared())))
}

func (v Vec3[T]) LenSquared() T {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v *Vec3[T]) Add(other Vec3[T]) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

func (v Vec3[T]) Equals(other Vec3[T]) bool {
	return v.X == other.X &&
		v.Y == other.Y &&
		v.Z == other.Z
}
