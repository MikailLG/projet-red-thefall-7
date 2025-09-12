package character

import (
	"fmt"

	"github.com/MikailLG/projet-red-thefall-7/src/character"
)

func DisplayInfo(personnage character.Character) character.Character {
	fmt.Println("===Informations du personnage===")
	fmt.Println("Nom:", personnage.Nom)
	fmt.Println("Classe:", personnage.Classe)
	fmt.Println("Niveau:", personnage.Niveau)
	fmt.Println("PointDeVie", personnage.PointDeVie)
	fmt.Println("PointDeVieMax", personnage.PointDeVieMax)
	fmt.Println("Inventaire:", personnage.Inventaire)
	return personnage
}
