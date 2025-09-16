package character

import (
	"fmt"
)

func apprendreCompétence(p *Character, apprendre string) {
	for _, s := range p.compétence {
		if s == apprendre {
			fmt.Println("Vous avez déja appris cette compétence :", apprendre)
			return
		}
	}
	p.compétence = append(p.compétence, apprendre)
	fmt.Println("Nouvelle compétence apprise :", apprendre)
}
