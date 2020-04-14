package resolvers

import (
	"errors"

	"github.com/graphql-go/graphql/gqlerrors"
)

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

func GetHuman(id int) (StarWarsChar, error) {
	for _, human := range HumanData {
		if human.ID == id {
			return human, nil
		}
	}
	return StarWarsChar{}, gqlerrors.FormatError(errors.New("cound not find the person with given id"))
}
