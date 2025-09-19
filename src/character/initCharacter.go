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
	Inventaire    map[string]int
	Competences   []string
	Argent        int
	Equipement    Equipement
}

type Equipement struct {
	Casque        string
	GiletParBalle string
	Botte         string
}

type EquipementCharacter struct {
	Nom        string
	PointDeVie int
	Equipement Equipement
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
			fmt.Printf("%sErreur : classe inconnue. Veuillez entrer Militaire, Citoyen ou Braconnier.%s", Red, Reset)
			continue
		}
		break
	}

	return Character{
		Nom:           "",
		Classe:        classe,
		Niveau:        1,
		PointDeVieMax: pvMax,
		PointDeVie:    pvMax,
		Inventaire: map[string]int{
			"Lampe de poche": 1,
			"Couteau suisse": 1,
		},
		Competences: []string{""},
		Argent:      100,
		Equipement: Equipement{
			Casque:        "Aucun",
			GiletParBalle: "Aucun",
			Botte:         "Aucune",
		},
	}
}
