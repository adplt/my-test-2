package main

import (
	"encoding/json"
	"fmt"
	configs "project/app/configs"
	structs "project/app/structs"
)

func main() {
	configs.Block{
		Try: func() {
			/* Manipulate Motorboat */
			motorboat1 := &structs.Motorboat{
				Ship: structs.Ship{
					Name:         "Bertram",
					Length:       10.66,
					Speed:        39,
					FuelCapacity: 310,
				},
			}
			motorboat2 := &structs.Motorboat{
				Ship: structs.Ship{
					Name:         "Boston Whaler",
					Length:       5.28,
					Speed:        33,
					FuelCapacity: 25,
				},
			}
			motorboat3 := &structs.Motorboat{
				Ship: structs.Ship{
					Name:         "Chaparral",
					Length:       2.54,
					Speed:        35,
					FuelCapacity: 40,
				},
			}

			motorBoats := [...]structs.Motorboats{motorboat1, motorboat2, motorboat3}
			for _, m := range motorBoats {
				m.SaveMotorboat()
				motorboat := m.GetMotorboat()
				motorboatJson, err := json.Marshal(motorboat)
				if err != nil {
					configs.Throw(err)
				}
				fmt.Println("Get Motorboat: ", string(motorboatJson))
			}

			/* **************************************************************************** */

			/* Manipulate Sailboat */
			sailboat1 := &structs.Sailboat{
				Ship: structs.Ship{
					Name:         "Bertram",
					Length:       10.66,
					Speed:        39,
					FuelCapacity: 310,
				},
			}
			sailboat2 := &structs.Sailboat{
				Ship: structs.Ship{
					Name:         "Boston Whaler",
					Length:       5.28,
					Speed:        33,
					FuelCapacity: 25,
				},
			}
			sailboat3 := &structs.Sailboat{
				Ship: structs.Ship{
					Name:         "Chaparral",
					Length:       2.54,
					Speed:        35,
					FuelCapacity: 40,
				},
			}

			sailBoats := [...]structs.Sailboats{sailboat1, sailboat2, sailboat3}
			for _, m := range sailBoats {
				m.SaveSailboat()
				sailboat := m.GetSailboat()
				sailboatJson, err := json.Marshal(sailboat)
				if err != nil {
					configs.Throw(err)
				}
				fmt.Println("Get Sailboat: ", string(sailboatJson))
			}

			/* **************************************************************************** */

			/* Manipulate Cruise */
			cruise1 := &structs.Cruise{
				Ship: structs.Ship{
					Name:         "Bertram",
					Length:       10.66,
					Speed:        39,
					FuelCapacity: 310,
				},
			}
			cruise2 := &structs.Cruise{
				Ship: structs.Ship{
					Name:         "Boston Whaler",
					Length:       5.28,
					Speed:        33,
					FuelCapacity: 25,
				},
			}
			cruise3 := &structs.Cruise{
				Ship: structs.Ship{
					Name:         "Chaparral",
					Length:       2.54,
					Speed:        35,
					FuelCapacity: 40,
				},
			}

			cruises := [...]structs.Cruises{cruise1, cruise2, cruise3}
			for _, m := range cruises {
				m.SaveCruise()
				cruise := m.GetCruise()
				cruiseJson, err := json.Marshal(cruise)
				if err != nil {
					configs.Throw(err)
				}
				fmt.Println("Get Cruise: ", string(cruiseJson))
			}
		},
		Catch: func(e error) {
			fmt.Println(e)
		},
	}.Do()
}
