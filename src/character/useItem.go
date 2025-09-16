package character

import (
	"fmt"
)

func UseItem(p *Character, item string) {
	for i, c := range p.Inventaire {
		if c == item {
			if c == "Compétence : Technique de combat avancée" {
				apprendreCompétence(p, "Technique de combat avancée")
			} else if c == "Bandage" {
				p.PointDeVie += 15
				if p.PointDeVie > p.PointDeVieMax {
					p.PointDeVie = p.PointDeVieMax
				}
				fmt.Printf("Vous utilisez un bandage.\nPV : %d / %d\n", p.PointDeVie, p.PointDeVieMax)
			} else {
				fmt.Println("Vous utilisez :", c)
			}

			p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			return
		}
	}
	fmt.Println("Objet non trouvé dans l’inventaire :", item)
}
