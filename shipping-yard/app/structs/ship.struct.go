package structs

import (
	"fmt"
)

type Ship struct {
	Name           string
	Length         float32
	Speed          int
	FuelCapacity   int
	ShieldStrength int
}

type Motorboat struct {
	Ship
}

type Motorboats interface {
	SaveMotorboat()
	GetMotorboat() Motorboat
}

func (s Motorboat) SaveMotorboat() {
	fmt.Println("Save Motorboat: ", s.Name, s.Length, s.Speed, s.FuelCapacity, s.ShieldStrength)
}

func (s Motorboat) GetMotorboat() Motorboat {
	return s
}

type Sailboat struct {
	Ship
}

type Sailboats interface {
	SaveSailboat()
	GetSailboat() Sailboat
}

func (sb Sailboat) SaveSailboat() {
	fmt.Println("Save Sailboat: ", sb.Name, sb.Length, sb.Speed, sb.FuelCapacity, sb.ShieldStrength)
}

func (s Sailboat) GetSailboat() Sailboat {
	return s
}

type Cruise struct {
	Ship
}

type Cruises interface {
	SaveCruise()
	GetCruise() Cruise
}

func (c Cruise) SaveCruise() {
	fmt.Println("Save Cruise: ", c.Name, c.Length, c.Speed, c.FuelCapacity, c.ShieldStrength)
}

func (c Cruise) GetCruise() Cruise {
	return c
}
