package main

import "fmt"

type Observer interface {
    Update(data string)
}

type Subject interface {
    Register(o Observer)
    Unregister(o Observer)
    NotifyAll()
}

type WeatherStation struct {
    observers map[Observer]struct{}
    temperature string
}

func NewWeatherStation() *WeatherStation {
    return &WeatherStation{
        observers: make(map[Observer]struct{}),
    }
}

func (ws *WeatherStation) Register(o Observer) {
    ws.observers[o] = struct{}{}
}

func (ws *WeatherStation) Unregister(o Observer) {
    delete(ws.observers, o)
}

func (ws *WeatherStation) NotifyAll() {
    for o := range ws.observers {
        o.Update(ws.temperature)
    }
}

// Method to update state and notify observers
func (ws *WeatherStation) SetTemperature(temp string) {
    ws.temperature = temp
    ws.NotifyAll()
}

type PhoneDisplay struct {
    id string
}

func (p *PhoneDisplay) Update(data string) {
    fmt.Printf("PhoneDisplay %s updating temperature to %s\n", p.id, data)
}

type WindowDisplay struct {}

func (w *WindowDisplay) Update(data string) {
    fmt.Printf("WindowDisplay showing temperature %s\n", data)
}

func main() {
    ws := NewWeatherStation()

    phone1 := &PhoneDisplay{id: "A"}
    phone2 := &PhoneDisplay{id: "B"}
    window := &WindowDisplay{}

    ws.Register(phone1)
    ws.Register(phone2)
    ws.Register(window)

    ws.SetTemperature("25°C")

    // Unregister phone1 and update temperature
    ws.Unregister(phone1)
    ws.SetTemperature("30°C")
}


 
