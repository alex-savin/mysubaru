package main

import (
    "github.com/alex-savin/mysubaru"
)

func main() {
    car, _ := mysubaru.New()

    car.R().Unlock()
    car.R().Lock()

    car.R().GetLocation()
    car.R().GetHealth()
    car.R().GetClimateSettings()

    car.R().StartLightHorn()
    time.Sleep(5 * time.Second)
    car.R().StopLightHorn()

    car.R().StartEngine()
    time.Sleep(10 * time.Second)
    car.R().StopEngine()
}
