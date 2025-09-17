package character

import (
	"fmt"
	"strings"
)

type Character struct {
	Nom           string
	Classe        string
	Niveau        int
	PointDeVie    int
	PointDeVieMax int
	Inventaire    []string
	Competences   []string
}

func InitCharacter() Character {
	var classe string
	pvMax := 100

	for {
		fmt.Println("Choisissez votre classe : Militaire, Citoyen ou Braconnier")
		fmt.Scanln(&classe)
		classe = Capitalize(strings.TrimSpace(classe))

		switch classe {
		case "Militaire":
			pvMax = 120
		case "Citoyen":
			pvMax = 80
		case "Braconnier":
			pvMax = 100
		default:
			fmt.Println("Erreur : classe inconnue. Veuillez entrer Militaire, Citoyen ou Braconnier.")
			continue
		}
		break
	}
	return Character{
		Nom:           "Manu",
		Classe:        classe,
		Niveau:        1,
		PointDeVieMax: pvMax,
		PointDeVie:    pvMax,
		Inventaire:    []string{"Lampe de poche", "Couteau suisse"},
		Competences:   []string{"Coup de Poing"},
	}
}
