package resolvers

type StarWarsChar struct {
	ID              int
	Name            string
	Friends         []StarWarsChar
	AppearsIn       []int
	HomePlanet      string
	PrimaryFunction string
}

var HumanData = []StarWarsChar{}

func GetHuman(id int) StarWarsChar {
	for _, human := range HumanData {
		if human.ID == id {
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
