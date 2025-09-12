package character

type Character struct {
	Nom           string
	Classe        string
	Niveau        int
	PointDeVie    int
	PointDeVieMax int
	Inventaire    []string
}

func InitCharacter() Character {
	return Character{
		Nom:           "Manu",
		Classe:        "Humain",
		Niveau:        1,
		PointDeVie:    100,
		PointDeVieMax: 100,
		Inventaire:    []string{"Lampe de poche", "Couteau suissse"},
	}
}
