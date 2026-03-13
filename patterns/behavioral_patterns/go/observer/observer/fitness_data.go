package observer

type FitnessData struct {
	observers []FitnessDataObserver
	steps     int
	calories  int
	distance  float64
}

func NewFitnessData() *FitnessData {
	return &FitnessData{
		observers: make([]FitnessDataObserver, 0),
	}
}

func (f *FitnessData) RegisterObserver(observer FitnessDataObserver) {
	f.observers = append(f.observers, observer)
}

func (f *FitnessData) RemoveObserver(observer FitnessDataObserver) {

	for i, obs := range f.observers {
		if obs == observer {
			f.observers = append(f.observers[:i], f.observers[i+1:]...)
			break
		}
	}
}

func (f *FitnessData) NotifyObservers() {

	for _, obs := range f.observers {
		obs.Update(f.steps, f.calories, f.distance)
	}
}

func (f *FitnessData) SetMeasurments(steps int, calories int, distance float64) {
	f.steps = steps
	f.distance = distance
	f.calories = calories
	f.NotifyObservers()
}
