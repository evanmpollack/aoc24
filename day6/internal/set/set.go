package set

type Set[T comparable] struct {
	set map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{set: make(map[T]bool)}
}

func Add[T comparable](s *Set[T], value T) {
	s.set[value] = true
}

func Contains[T comparable](s *Set[T], value T) bool {
	_, ok := s.set[value]
	return ok
}
