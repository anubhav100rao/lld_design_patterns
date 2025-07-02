package observer

import "fmt"

// Observer interface
type PropertyObserver interface {
	OnChange(property string, oldValue, newValue interface{})
}

// Subject with observable properties
type Person struct {
	name      string
	age       int
	observers []PropertyObserver
}

func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func (p *Person) Register(o PropertyObserver) {
	p.observers = append(p.observers, o)
}

func (p *Person) notify(prop string, oldVal, newVal interface{}) {
	for _, o := range p.observers {
		o.OnChange(prop, oldVal, newVal)
	}
}

func (p *Person) SetName(name string) {
	old := p.name
	p.name = name
	p.notify("name", old, name)
}

func (p *Person) SetAge(age int) {
	old := p.age
	p.age = age
	p.notify("age", old, age)
}

// Concrete Observer
type Logger struct{}

func (l *Logger) OnChange(property string, oldValue, newValue interface{}) {
	fmt.Printf("Property %s changed from %v to %v\n", property, oldValue, newValue)
}

// Client
func RunPropertyChangeObserver() {
	person := NewPerson("Alice", 30)
	logger := &Logger{}

	person.Register(logger)

	person.SetName("Alicia")
	person.SetAge(31)
}
