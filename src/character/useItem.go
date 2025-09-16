package character

import (
	"fmt"
)

func UseItem(p *Character, item string) {
	for i, c := range p.Inventaire {
		if c == item {
			switch c {
			case "Compétence : Technique de combat avancée":
				apprendreCompétence(p, "Technique de combat avancée")
			case "Bandage":
				p.PointDeVie += 15
				if p.PointDeVie > p.PointDeVieMax {
					p.PointDeVie = p.PointDeVieMax
				}
				fmt.Printf("Vous utilisez un bandage.\nPV : %d / %d\n", p.PointDeVie, p.PointDeVieMax)
			default:
				fmt.Println("Vous utilisez :", c)
			}
			p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			return
		}
	}
	fmt.Println("Objet non trouvé dans l’inventaire :", item)
}
