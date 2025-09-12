package character

import (
	"fmt"
)

func DisplayInfo(personnage Character) Character {
	fmt.Println("===Informations du personnage===")
	fmt.Println("Nom:", personnage.Nom)
	fmt.Println("Classe:", personnage.Classe)
	fmt.Println("Niveau:", personnage.Niveau)
	fmt.Println("PointDeVie", personnage.PointDeVie)
	fmt.Println("PointDeVieMax", personnage.PointDeVieMax)
	fmt.Println("Inventaire:", personnage.Inventaire)
	return personnage
}
