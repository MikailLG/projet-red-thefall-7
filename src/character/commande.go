package character

import "fmt"

func DisplayInfo(p Character) {
	fmt.Println("==============================")
	fmt.Println("   Informations du personnage")
	fmt.Println("==============================")
	fmt.Printf("%sNom         : %s\n%s", Orange, p.Nom, Reset)
	fmt.Printf("%sClasse      : %s\n%s", Orange, p.Classe, Reset)
	fmt.Printf("%sNiveau      : %d\n%s", Orange, p.Niveau, Reset)
	fmt.Printf("%sPoints de vie : %d / %d\n%s", Orange, p.PointDeVie, p.PointDeVieMax, Reset)
	fmt.Printf("%sÉquipement  :%s\n", Orange, Reset)
	fmt.Printf("%sSlots d’inventaire : %d%s\n", Orange, p.CapaciteMax, Reset)
	fmt.Printf("  - Tête  : %s\n", p.Equipement.Casque)
	fmt.Printf("  - Torse : %s\n", p.Equipement.GiletParBalle)
	fmt.Printf("  - Pieds : %s\n", p.Equipement.Botte)

	fmt.Println(Orange, "Inventaire  :")
	for nom, quantite := range p.Inventaire {
		fmt.Printf("  - %s x%d\n", nom, quantite)
	}

	fmt.Println(Orange, "Compétences :")
	for _, skill := range p.Competences {
		fmt.Printf("  - %s\n", skill)
	}

	fmt.Printf("%sArgent : %d pièces\n%s", Yellow, p.Argent, Reset)
}
