package factory

type Shape interface {
	Draw() string
}

type Circle struct{}

func (c Circle) Draw() string {
	return "Drawing a Circle"
}

type Square struct{}

func (s Square) Draw() string {
	return "Drawing a Square"
}

type ShapeFactory struct{}

func (f ShapeFactory) CreateShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}

func RunFactory() {
	factory := ShapeFactory{}

	circle := factory.CreateShape("circle")
	if circle != nil {
		println(circle.Draw())
	}

	square := factory.CreateShape("square")
	if square != nil {
		println(square.Draw())
	}
}
