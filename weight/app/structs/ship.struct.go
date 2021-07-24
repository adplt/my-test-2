package structs

type Ship struct {
	Name           string
	Length         int
	Speed          int
	EnergyCapacity int
	ShieldStrength int
}

type Motorboats struct {
	Ship
}

type Sailboats struct {
	Ship
}

type Cruise struct {
	Ship
}
