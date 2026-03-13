package main

import "observer_method/observer"

func main() {
	fitnessData := observer.NewFitnessData()

	liveDisplay := observer.NewLiveActivityDisplay()
	logger := observer.NewProgressLogger()
	goalNotifier := observer.NewGoalNotifer(10000)

	fitnessData.RegisterObserver(liveDisplay)
	fitnessData.RegisterObserver(logger)
	fitnessData.RegisterObserver(goalNotifier)

	fitnessData.SetMeasurments(3000, 120, 2.5)
	fitnessData.SetMeasurments(7000, 250, 5.0)
	fitnessData.SetMeasurments(10000, 400, 8.0)

	fitnessData.RemoveObserver(logger)
	fitnessData.SetMeasurments(12000, 500, 10.0)
}
