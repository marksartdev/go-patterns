package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/observer"
)

func main() {
	wd := observer.NewWeatherData()
	currentConditionalsDisplay := observer.NewCurrentConditionsDisplay(wd)
	statisticsDisplay := observer.NewStatisticsDisplay(wd)
	forecastDisplay := observer.NewForecastDisplay(wd)
	heatIndexDisplay := observer.NewHeatIndexDisplay(wd)

	fmt.Println("Update 0 (without setChanged):")
	fmt.Println(wd.SetMeasurements(30, 60, 560))

	fmt.Println("Update 1:")
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(30, 60, 560))

	fmt.Println("Update 2:")
	wd.RemoveObserver(heatIndexDisplay)
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(25, 45, 600))

	fmt.Println("Update 3:")
	wd.RemoveObserver(forecastDisplay)
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(10, 50, 620))

	fmt.Println("Update 4:")
	wd.RemoveObserver(statisticsDisplay)
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(30, 80, 700))

	fmt.Println("Update 5:")
	wd.RemoveObserver(currentConditionalsDisplay)
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(0, 40, 680))
}
