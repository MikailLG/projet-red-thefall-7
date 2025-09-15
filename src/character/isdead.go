package character

import (
	"fmt"
)

func isDead(p *Character) bool {
	if p.HP <= 0 {
		fmt.Printf("%s est mort(e)...\n", &p.Name)
		p.HP = p.MaxHP / 2
		fmt.Printf("%s ressuscite avec %d PV.\n", &p.Name, p.HP)
		return true
	}
	return false
}
