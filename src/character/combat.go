package character

import (
	"fmt"
	"time"
)

func apprendreCompetence(p *Character, apprendre string) {
	for _, s := range p.Competences {
		if s == apprendre {
			fmt.Printf("%sVous avez déjà appris :%s", Red, apprendre)
			return
		}
	}
	p.Competences = append(p.Competences, apprendre)
	fmt.Printf("%sNouvelle compétence apprise :%s", Green, apprendre)
}

func Heal(p *Character) {
	if RemoveItem(p, "Bandage") {
		p.PointDeVie += 15
		if p.PointDeVie > p.PointDeVieMax {
			p.PointDeVie = p.PointDeVieMax
		}
		fmt.Printf("%sVous utilisez un bandage.\nPV : %d / %d\n%s", Green, p.PointDeVie, p.PointDeVieMax, Reset)
	} else {
		fmt.Printf("%sVous n'avez pas de bandage dans votre inventaire.%s", Red, Reset)
	}
}

func UseItem(p *Character, item string) {
	quantite, exists := p.Inventaire[item]
	if !exists {
		fmt.Printf("%sObjet non trouvé dans l’inventaire :%s", item, Red)
		return
	}

	switch item {
	case "Technique de combat avancée":
		apprendreCompetence(p, "Technique de combat avancée")

	case "Bandage":
		p.PointDeVie += 15
		if p.PointDeVie > p.PointDeVieMax {
			p.PointDeVie = p.PointDeVieMax
		}
		fmt.Printf("%sVous utilisez un bandage.\nPV : %d / %d\n%s", Green, p.PointDeVie, p.PointDeVieMax, Reset)
	case "Casque renforcée":
		if p.Equipement.Casque == "Aucun" {
			p.Equipement.Casque = "Casque renforcée"
			p.PointDeVieMax += 15
			p.PointDeVie += 15
			fmt.Printf("%sVous équipez un Casque. +15 PV max.%s", Green, Reset)
		} else {
			fmt.Printf("%sVous portez déjà un casque.%s", Red, Reset)
			return
		}
	case "Gilet pare-balles":
		if p.Equipement.GiletParBalle == "Aucun" {
			p.Equipement.GiletParBalle = "Gilet pare-balles"
			p.PointDeVieMax += 25
			p.PointDeVie += 25
			fmt.Printf("%sVous équipez un Gilet pare-balles. +25 PV max.%s", Green, Reset)
		} else {
			fmt.Printf("%sVous portez déjà un gilet.%s", Red, Reset)
			return
		}
	case "Bottes renforcée":
		if p.Equipement.Botte == "Aucune" {
			p.Equipement.Botte = "Bottes renforcée"
			p.PointDeVieMax += 10
			p.PointDeVie += 10
			fmt.Printf("%sVous équipez des Bottes renforcée. +10 PV max.%s", Green, Reset)
		} else {
			fmt.Printf("%sVous portez déjà des bottes.%s", Red, Reset)
			return
		}

	default:
		fmt.Println("Vous utilisez :", item)
	}
	if quantite > 1 {
		p.Inventaire[item]--
	} else {
		delete(p.Inventaire, item)
	}
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
