package lib

/**

Example of writing extensible library with unsettable options

*/

// Car object
type Car struct {
	Wheels int
}

// NewCar returns default Car
func NewCar() *Car {
	return &Car{4}
}

// Option is a way of modifiying internal fields, it returns a function which allows to unset the last applied option
func (c *Car) Option(opts ...CarOption) (previous CarOption) {
	for _, opt := range opts {
		previous = opt(c)
	}
	return previous
}

// CarOption allows extending default Car with possibility of unsetting values
type CarOption func(*Car) CarOption

// Wheels allows to temporarily change the number of Wheels
func Wheels(numWheels int) CarOption {
	return func(c *Car) CarOption {
		curWheels := c.Wheels
		c.Wheels = numWheels
		return Wheels(curWheels)
	}
}
