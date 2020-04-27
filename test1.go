package main

import (
	EventCounter "event_counter/event-counter/event-counter"
	"fmt"
	"time"
)

func main() {
	overRpmCount := new(EventCounter.EventCounter)
	overRpmCount.Name = "over rpm count"
	overRpmCount.Value = 50
	overRpmCount.LastUpdate = time.Now()
	fmt.Printf("%+v\n", overRpmCount.Name)
	fmt.Printf("%+v\n", overRpmCount.LastUpdate)

	overSpeedCount := new(EventCounter.EventCounter)
	overSpeedCount.Name = "over speed count"
	overSpeedCount.Value = 500
	overSpeedCount.LastUpdate = time.Now()
	fmt.Printf("%+v\n", overSpeedCount.Name)

	overRpmCount.Inc()
	fmt.Println(overRpmCount.Value)
	fmt.Println(overRpmCount.LastUpdate)

	overSpeedCount.Inc()
	fmt.Println(overSpeedCount.Value)
}
