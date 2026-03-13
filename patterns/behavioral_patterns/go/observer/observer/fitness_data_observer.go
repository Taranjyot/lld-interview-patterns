package observer

type FitnessDataObserver interface {
	Update(steps int, calories int, distance float64)
}
