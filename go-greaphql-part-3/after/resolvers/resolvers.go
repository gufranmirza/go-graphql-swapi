package resolvers

import "fmt"

type StarWarsChar struct {
	ID              int
	Name            string
	Friends         []StarWarsChar
	AppearsIn       []int
	HomePlanet      string
	PrimaryFunction string
}

var Luke = StarWarsChar{
	ID:         1000,
	Name:       "Luke Skywalker",
	AppearsIn:  []int{4, 5, 6},
	HomePlanet: "Tatooine",
}

var Vader = StarWarsChar{
	ID:         1001,
	Name:       "Darth Vader",
	AppearsIn:  []int{4, 5, 6},
	HomePlanet: "Tatooine",
}

var Han = StarWarsChar{
	ID:         1002,
	Name:       "Han Solo",
	HomePlanet: "Alderaa",
	AppearsIn:  []int{4, 5, 6},
}

var Leia = StarWarsChar{
	ID:         1003,
	Name:       "Leia Organa",
	AppearsIn:  []int{4, 5, 6},
	HomePlanet: "Alderaa",
}

var Tarkin = StarWarsChar{
	ID:         1004,
	Name:       "Wilhuff Tarkin",
	AppearsIn:  []int{4},
	HomePlanet: "Alderaa",
}
var Threepio = StarWarsChar{
	ID:              2000,
	Name:            "C-3PO",
	AppearsIn:       []int{4, 5, 6},
	PrimaryFunction: "Protocol",
}
var Artoo = StarWarsChar{
	ID:              2001,
	Name:            "R2-D2",
	AppearsIn:       []int{4, 5, 6},
	PrimaryFunction: "Astromech",
}

var HumanData = []StarWarsChar{
	Luke,
	Vader,
	Han,
	Leia,
	Tarkin,
	Threepio,
	Artoo,
}

func GetHuman(id int) StarWarsChar {
	for _, human := range HumanData {
		if human.ID == id {
			return human
		}
	}
	return StarWarsChar{}
}

func GetHumans() []StarWarsChar {
	return HumanData
}

func UpdatePerson(id int, name string, appearsIn []int, homeplanet string) StarWarsChar {
	for _, human := range HumanData {
		fmt.Println(human, id)
		if human.ID == id {
			human.Name = name
			human.AppearsIn = appearsIn
			human.HomePlanet = homeplanet
			return human
		}
	}

	return StarWarsChar{}
}

func CreatePerson(id int, name string, appearsIn []int, homeplanet string) StarWarsChar {
	char := StarWarsChar{
		ID:         id,
		Name:       name,
		AppearsIn:  appearsIn,
		HomePlanet: homeplanet,
	}

	HumanData = append(HumanData, char)

	return char
}
