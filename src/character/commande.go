package character

import "fmt"

func DisplayInfo(p Character) {
	fmt.Println("==============================")
	fmt.Println("   Informations du personnage")
	fmt.Println("==============================")
	fmt.Printf("Nom         : %s\n", p.Nom)
	fmt.Printf("Classe      : %s\n", p.Classe)
	fmt.Printf("Niveau      : %d\n", p.Niveau)
	fmt.Printf("Points de vie : %d / %d\n", p.PointDeVie, p.PointDeVieMax)

	fmt.Println("Inventaire  :")
	for nom, quantite := range p.Inventaire {
		fmt.Printf("  - %s x%d\n", nom, quantite)
	}
	fmt.Println("Compétences :")
	for _, skill := range p.Competences {
		fmt.Printf("  - %s\n", skill)
	}
	fmt.Printf("Argent : %d pièces\n", p.Argent)
}
