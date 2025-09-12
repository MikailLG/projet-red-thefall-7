package character

import "fmt"

func Heal(p *Character) {
	for i, item := range p.Inventaire {
		if item == "Bandage" {
			p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			p.PointDeVie += 15
			if p.PointDeVie > p.PointDeVieMax {
				p.PointDeVie = p.PointDeVieMax
			}

			fmt.Printf("Vous utilisez un bandage.\nPV : %d / %d\n", p.PointDeVie, p.PointDeVieMax)
			return
		}
	}
	fmt.Println("Vous n'avez pas de bandage dans votre inventaire.")
}
