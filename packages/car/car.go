package car

type Car struct {
	Name  string
	Color string
}

func (c Car) Start() string {
	return "The car " + c.Name + " which has a " + c.Color + " color has been started"
}
