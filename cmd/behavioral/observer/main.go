package main

import (
	"fmt"
	"math/rand"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/observer"
)

const (
	minTemperature = 25
	maxTemperature = 40
	minHumidity    = 40
	maxHumidity    = 80
	minPressure    = 560
	maxPressure    = 700
)

func main() {
	var (
		currentConditionalsDisplay observer.DisplayElement
		statisticsDisplay          observer.DisplayElement
		forecastDisplay            observer.DisplayElement
		heatIndexDisplay           observer.DisplayElement
		wd                         observer.WeatherDater
	)

	wd = observer.NewWeatherData()
	currentConditionalsDisplay = observer.NewCurrentConditionsDisplay(wd)
	statisticsDisplay = observer.NewStatisticsDisplay(wd)
	forecastDisplay = observer.NewForecastDisplay(wd)
	heatIndexDisplay = observer.NewHeatIndexDisplay(wd)

	fmt.Println("Update 0 (without setChanged):")
	fmt.Println(wd.SetMeasurements(getTemperature(), getHumidity(), getPressure()))

	fmt.Println("Update 1:")
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(getTemperature(), getHumidity(), getPressure()))

	fmt.Println("Update 2:")
	wd.RemoveObserver(heatIndexDisplay)
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(getTemperature(), getHumidity(), getPressure()))

	fmt.Println("Update 3:")
	wd.RemoveObserver(forecastDisplay)
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(getTemperature(), getHumidity(), getPressure()))

	fmt.Println("Update 4:")
	wd.RemoveObserver(statisticsDisplay)
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(getTemperature(), getHumidity(), getPressure()))

	fmt.Println("Update 5:")
	wd.RemoveObserver(currentConditionalsDisplay)
	wd.SetChanged()
	fmt.Println(wd.SetMeasurements(getTemperature(), getHumidity(), getPressure()))
}

func getTemperature() float64 {
	return getRand(minTemperature, maxTemperature)
}

func getHumidity() float64 {
	return getRand(minHumidity, maxHumidity)
}

func getPressure() float64 {
	return getRand(minPressure, maxPressure)
}

func getRand(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
