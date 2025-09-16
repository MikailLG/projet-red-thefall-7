package character

import "fmt"

func DisplayInfo(p Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom:", p.Nom)
	fmt.Println("Classe:", p.Classe)
	fmt.Println("Niveau:", p.Niveau)
	fmt.Println("PointDeVie:", p.PointDeVie)
	fmt.Println("PointDeVieMax:", p.PointDeVieMax)
	fmt.Println("Inventaire:", p.Inventaire)
	fmt.Println("Compétences:", p.Competences)
}
