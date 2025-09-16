package character

import (
	"fmt"
	"time"
)

func apprendreCompetence(p *Character, apprendre string) {
	for _, s := range p.Competences {
		if s == apprendre {
			fmt.Println("Vous avez déjà appris :", apprendre)
			return
		}
	}
	p.Competences = append(p.Competences, apprendre)
	fmt.Println("Nouvelle compétence apprise :", apprendre)
}

func Heal(p *Character) {
	if RemoveItem(p, "Bandage") {
		p.PointDeVie += 15
		if p.PointDeVie > p.PointDeVieMax {
			p.PointDeVie = p.PointDeVieMax
		}
		fmt.Printf("Vous utilisez un bandage.\nPV : %d / %d\n", p.PointDeVie, p.PointDeVieMax)
	} else {
		fmt.Println("Vous n'avez pas de bandage dans votre inventaire.")
	}
}

func UseItem(p *Character, item string) {
	for i, c := range p.Inventaire {
		if c == item {
			switch c {
			case "Compétence : Technique de combat avancée":
				apprendreCompetence(p, "Technique de combat avancée")
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

func IsDead(p *Character) bool {
	if p.PointDeVie <= 0 {
		fmt.Printf("%v est mort(e)...\n", p.Nom)
		p.PointDeVie = p.PointDeVieMax / 2
		fmt.Printf("%v ressuscite avec %d PV.\n", p.Nom, p.PointDeVie)
		return true
	}
	return false
}

func Poison(p *Character) {
	for i := 1; i <= 3; i++ {
		p.PointDeVie -= 10
		if p.PointDeVie < 0 {
			p.PointDeVie = 0
		}
		fmt.Printf("PV actuels : %d/%d\n", p.PointDeVie, p.PointDeVieMax)
		time.Sleep(1 * time.Second)
	}
}
