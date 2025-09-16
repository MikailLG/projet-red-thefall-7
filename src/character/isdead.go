package character

import (
	"fmt"
)

func IsDead(p *Character) bool {
	if p.PointDeVie <= 0 {
		fmt.Printf("%v est mort(e)...\n", &p.Nom)
		p.PointDeVie = p.PointDeVieMax / 2
		fmt.Printf("%v ressuscite avec %d PV.\n", &p.Nom, p.PointDeVie)
		return true
	}
	return false
}
