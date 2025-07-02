package observer

import "fmt"

// Observer interface
type DisplayObserver interface {
	Update(temp, humidity float64)
}

// Subject interface
type WeatherSubject interface {
	Register(o DisplayObserver)
	Unregister(o DisplayObserver)
	Notify()
}

// ConcreteSubject
type WeatherStation struct {
	observers   map[DisplayObserver]struct{}
	temperature float64
	humidity    float64
}

func NewWeatherStation() *WeatherStation {
	return &WeatherStation{observers: make(map[DisplayObserver]struct{})}
}

func (ws *WeatherStation) Register(o DisplayObserver) {
	ws.observers[o] = struct{}{}
}

func (ws *WeatherStation) Unregister(o DisplayObserver) {
	delete(ws.observers, o)
}

func (ws *WeatherStation) Notify() {
	for o := range ws.observers {
		o.Update(ws.temperature, ws.humidity)
	}
}

func (ws *WeatherStation) MeasurementsChanged() {
	ws.Notify()
}

func (ws *WeatherStation) SetMeasurements(temp, humidity float64) {
	ws.temperature = temp
	ws.humidity = humidity
	ws.MeasurementsChanged()
}

// ConcreteObservers
type CurrentConditionsDisplay struct {
	name string
}

func NewCurrentDisplay(name string) *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{name}
}

func (d *CurrentConditionsDisplay) Update(temp, hum float64) {
	fmt.Printf("[%s] Current conditions: %.1f°C, %.1f%% humidity\n",
		d.name, temp, hum)
}

type StatisticsDisplay struct{}

func (s *StatisticsDisplay) Update(temp, hum float64) {
	// Could track min/max/avg; here we just print.
	fmt.Printf("[Stats] Received temp=%.1f, humidity=%.1f\n", temp, hum)
}

// Client
func RunWeatherStationObserver() {
	station := NewWeatherStation()
	display1 := NewCurrentDisplay("LCD1")
	stats := &StatisticsDisplay{}

	station.Register(display1)
	station.Register(stats)

	station.SetMeasurements(22.5, 65.0)
	station.SetMeasurements(23.0, 70.0)

	station.Unregister(display1)
	station.SetMeasurements(21.0, 60.0)
}
