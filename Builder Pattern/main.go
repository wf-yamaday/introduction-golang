package main

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "grenn"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	// Sheet(int) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type carBuilder struct {
	color  Color
	wheels Wheels
	speed  Speed
}

type car struct {
	params carBuilder
}

func NewBuilder() *carBuilder {
	return &carBuilder{
		color:  BlueColor,
		wheels: SportsWheels,
		speed:  MPH,
	}
}

func (b *carBuilder) Color(color Color) Builder {
	b.color = color
	return b
}

func (b *carBuilder) Wheels(wheels Wheels) Builder {
	b.wheels = wheels
	return b
}

func (b *carBuilder) TopSpeed(speed Speed) Builder {
	b.speed = speed
	return b
}

func (b *carBuilder) Build() Interface {
	if b.Wheels == SportsWheels {
		return &sportsWheelsCar{
			params: *b,
		}
	}
	return &steelWheelsCar{
		params: *b,
	}
}

func (c *car) Drive() error {
	fmt.Println("Driving' %#=v\n", c.params)
	return nil
}

func (c *car) Stop() error {
	fmt.Println("Stop: %#+v\n", c.params)
	return nil
}

func main() {
	car := NewBuilder().Color(BlueColor).Wheels(SteelWheels).TopSpeed(KPH).Build()
	car.Drive()
	car.Stop()
}
