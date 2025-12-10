package utilities

type Rectangle[T Number] struct {
	Min Vec2[T]
	Max Vec2[T]
}

func (r Rectangle[T]) Canonical() Rectangle[T] {
	if r.Max.X < r.Min.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}
	if r.Max.Y < r.Min.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
	return r
}

func (r Rectangle[T]) Width() T {
	return r.Max.X - r.Min.X
}

func (r Rectangle[T]) Height() T {
	return r.Max.Y - r.Min.Y
}

func (r Rectangle[T]) Inset(n T) Rectangle[T] {
	if r.Width() < 2*n {
		r.Min.X = (r.Min.X + r.Max.X) / 2
		r.Max.X = r.Min.X
	} else {
		r.Min.X += n
		r.Max.X -= n
	}
	if r.Height() < 2*n {
		r.Min.Y = (r.Min.Y + r.Max.Y) / 2
		r.Max.Y = r.Min.Y
	} else {
		r.Min.Y += n
		r.Max.Y -= n
	}
	return r
}

func (r Rectangle[T]) IsEmpty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

func (r Rectangle[T]) Overlaps(s Rectangle[T]) bool {
	return !r.IsEmpty() && !s.IsEmpty() &&
		r.Min.X < s.Max.X && s.Min.X < r.Max.X &&
		r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}
