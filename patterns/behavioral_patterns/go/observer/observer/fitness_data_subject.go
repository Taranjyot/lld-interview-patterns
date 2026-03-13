package observer

type FitnessDataSubject interface {
	RegisterObserver(observer FitnessDataObserver)
	RemoveObserver(observer FitnessDataObserver)
	NotifyObservers()
}
